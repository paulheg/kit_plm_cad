import 'package:flutter/material.dart';
import 'package:flutter_joystick/flutter_joystick.dart';
import 'package:go_router/go_router.dart';
import 'package:kerberos_control/api/connection_service.dart';
import 'package:kerberos_control/settings.dart';
import 'package:kerberos_control/widgets/vision_controls.dart';
import 'package:provider/provider.dart';

class ControlPage extends StatefulWidget {
  final String robotId;

  const ControlPage(this.robotId, {super.key});

  @override
  State<ControlPage> createState() => _ControlPageState();
}

class _ControlPageState extends State<ControlPage> {
  double _speedX = 0;
  double _speedY = 0;

  final ConnectionService _connection = ConnectionService();

  @override
  void dispose() {
    _connection.disconnect();

    super.dispose();
  }

  void _showInSnackbar(String message) {
    if (context.mounted) {
      ScaffoldMessenger.of(context)
          .showSnackBar(SnackBar(content: Text(message)));
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Control - ${widget.robotId}"),
        leading: IconButton(
          icon: const Icon(Icons.arrow_back),
          onPressed: () {
            _connection.disconnect();
            context.go('/select');
          },
        ),
      ),
      body: Consumer<Settings>(
        builder: (context, settings, child) {
          return FutureBuilder<ConnectionService>(
            future: _connection.connect(settings.host, widget.robotId, () {
              _showInSnackbar("Connection was closed.");
              _connection.disconnect();
              context.go('/select');
            }),
            builder: (context, snapshot) {
              if (snapshot.hasData) {
                return Flex(
                    direction: Axis.vertical,
                    mainAxisAlignment: MainAxisAlignment.center,
                    crossAxisAlignment: CrossAxisAlignment.stretch,
                    verticalDirection: VerticalDirection.down,
                    children: [
                      Flex(
                        direction: MediaQuery.of(context).size.width >= 700
                            ? Axis.horizontal
                            : Axis.vertical,
                        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                        children: [
                          Padding(
                            padding: const EdgeInsets.all(10.0),
                            child: VisionControls(snapshot.data!),
                          ),
                          Padding(
                            padding: const EdgeInsets.all(20.0),
                            child: Joystick(
                              listener: (details) {
                                _updateStick(details);
                                _connection.send("move|$_speedX|$_speedY");
                              },
                            ),
                          ),
                        ],
                      ),
                    ]);
              } else if (snapshot.hasError) {
                return const Center(child: Text("Couldn't connect."));
              }

              return const Center(child: Text("Connecting..."));
            },
          );
        },
      ),
    );
  }

  void _updateStick(StickDragDetails details) {
    setState(() {
      _speedX = details.x;
      _speedY = details.y;
    });
  }
}
