import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:kerberos_control/api/calls.dart';
import 'package:kerberos_control/api/connection_status.dart';
import 'package:kerberos_control/settings.dart';
import 'package:provider/provider.dart';

class SelectRobotPage extends StatefulWidget {
  const SelectRobotPage({super.key});

  @override
  State<SelectRobotPage> createState() => _SelectRobotPageState();
}

class _SelectRobotPageState extends State<SelectRobotPage> {
  @override
  Widget build(BuildContext context) {
    List<ConnectionStatus> _connections = [];

    Future<void> refresh(String host) async {
      var newConns = await getConnectionStatusList(host);

      setState(() {
        _connections = newConns;
      });
    }

    Widget _getListView(Settings settings) {
      return RefreshIndicator(
        onRefresh: () => refresh(settings.host),
        child: ListView.builder(
          itemCount: _connections.length,
          itemBuilder: (context, index) {
            var conn = _connections[index];

            return GestureDetector(
              onTap: () {
                context.goNamed('control', pathParameters: {
                  'id': conn.id,
                });
              },
              child: Card(
                child: Padding(
                  padding: const EdgeInsets.all(20.0),
                  child: Flex(
                    direction: Axis.horizontal,
                    children: [
                      Padding(
                        padding: const EdgeInsets.only(right: 20),
                        child: SizedBox(
                            height: 60,
                            width: 60,
                            child: Image.asset("assets/Controller Logo.png")),
                      ),
                      Column(
                          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(
                              conn.id,
                              style:
                                  const TextStyle(fontWeight: FontWeight.bold),
                            ),
                            Flex(
                              direction: Axis.horizontal,
                              children: [
                                const Text("Open: "),
                                (conn.connected)
                                    ? const Icon(Icons.remove_circle_outline)
                                    : const Icon(Icons.check_circle_outline)
                              ],
                            )
                          ]),
                    ],
                  ),
                ),
              ),
            );
          },
        ),
      );
    }

    Widget _getMessageCard(String message) {
      return Center(
        child: Padding(
          padding: const EdgeInsets.all(20.0),
          child: Card(
            child: Center(
              child: Text(message),
            ),
          ),
        ),
      );
    }

    return Consumer<Settings>(
      builder: (context, settings, child) {
        return Scaffold(
          appBar: AppBar(
            title: const Text("Select Robot"),
            leading: IconButton(
              icon: const Icon(Icons.arrow_back),
              onPressed: () {
                context.go('/connect');
              },
            ),
            actions: [
              IconButton(
                  onPressed: () {
                    refresh(settings.host);
                  },
                  icon: const Icon(Icons.refresh))
            ],
          ),
          body: FutureBuilder<List<ConnectionStatus>>(
            future: getConnectionStatusList(settings.host),
            builder: (context, snapshot) {
              if (snapshot.hasData) {
                _connections = snapshot.data!;

                return _getListView(settings);
              } else if (snapshot.hasError) {
                _getMessageCard("Error while loading robots");
              }

              return _getMessageCard("Loading");
            },
          ),
        );
      },
    );
  }
}
