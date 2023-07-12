import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:kerberos_control/api/calls.dart';
import 'package:kerberos_control/settings.dart';
import 'package:provider/provider.dart';

class ConnectPage extends StatefulWidget {
  const ConnectPage({super.key});

  @override
  State<ConnectPage> createState() => _ConnectPageState();
}

class _ConnectPageState extends State<ConnectPage> {
  final _formKey = GlobalKey<FormState>();

  void _showInSnackbar(String message) {
    if (context.mounted) {
      ScaffoldMessenger.of(context)
          .showSnackBar(SnackBar(content: Text(message)));
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(title: const Text("Connect to server")),
        body: Consumer<Settings>(
          builder: (context, settings, child) {
            return FutureBuilder(
              future: settings.load(),
              builder: (context, snapshot) {
                if (snapshot.hasData) {
                  return Center(
                    child: Padding(
                      padding: const EdgeInsets.all(20.0),
                      child: Card(
                        child: Padding(
                          padding: const EdgeInsets.all(10.0),
                          child: Form(
                            key: _formKey,
                            child: Wrap(
                              children: [
                                const Text("Server"),
                                TextFormField(
                                  initialValue: settings.host,
                                  onSaved: (newValue) {
                                    if (newValue != null) {
                                      settings.setHost(newValue);
                                    }
                                  },
                                  validator: (value) {
                                    if (value == null || value.isEmpty) {
                                      return 'Please enter a hostname';
                                    }

                                    if (Uri.tryParse(value) == null) {
                                      return 'Please enter a valid hostname';
                                    }

                                    return null;
                                  },
                                ),
                                Center(
                                  child: Padding(
                                    padding: const EdgeInsets.all(15.0),
                                    child: ElevatedButton(
                                        onPressed: () async {
                                          if (_formKey.currentState!
                                              .validate()) {
                                            _formKey.currentState!.save();
                                            await settings.save();

                                            _showInSnackbar("Connecting...");

                                            try {
                                              var statusList =
                                                  await getConnectionStatusList(
                                                      settings.host);

                                              _showInSnackbar(
                                                  "Connected with ${statusList.length} robots.");

                                              context.go('/select');
                                            } catch (e) {
                                              _showInSnackbar(
                                                  "Connection failed: ${e.toString()}");
                                            }
                                          }
                                        },
                                        child: const Text("Connect")),
                                  ),
                                )
                              ],
                            ),
                          ),
                        ),
                      ),
                    ),
                  );
                }

                return const Center(
                  child: Text("Loading Settings"),
                );
              },
            );
          },
        ));
  }
}
