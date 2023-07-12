import 'package:flutter/material.dart';
import 'package:kerberos_control/api/connection_service.dart';
import 'package:kerberos_control/widgets/repeat_on_tap.dart';

class VisionControls extends StatefulWidget {
  final ConnectionService service;

  VisionControls(this.service, {super.key});

  @override
  State<VisionControls> createState() => _VisionControlsState();
}

class _VisionControlsState extends State<VisionControls> {
  double _height = 0;
  double _tilt = 0;
  bool _autoTilt = true;

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            const Text("Height"),
            Flex(
              direction: Axis.horizontal,
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                Padding(
                  padding: const EdgeInsets.all(8.0),
                  child: RepeatOnTap(
                    call: () {
                      widget.service.send("move_up");
                    },
                    child: ElevatedButton(
                        onPressed: () {},
                        child: const Padding(
                          padding: EdgeInsets.all(8.0),
                          child: Icon(Icons.arrow_circle_up),
                        )),
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.all(8.0),
                  child: RepeatOnTap(
                    call: () {
                      widget.service.send("move_down");
                    },
                    child: ElevatedButton(
                        onPressed: () {},
                        child: const Padding(
                          padding: EdgeInsets.all(8.0),
                          child: Icon(Icons.arrow_circle_down),
                        )),
                  ),
                )
              ],
            )
          ],
        ),
        Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            const Text("Auto Tilt"),
            Row(
              children: [
                Switch(
                  value: _autoTilt,
                  onChanged: (value) {
                    setState(() {
                      widget.service.send("auto_tilt|$value");
                      _autoTilt = value;
                    });
                  },
                ),
                Flex(
                  direction: Axis.horizontal,
                  mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                  children: [
                    Padding(
                      padding: const EdgeInsets.all(8.0),
                      child: RepeatOnTap(
                        call: () {
                          widget.service.send("tilt_up");
                        },
                        child: ElevatedButton(
                            onPressed: () {},
                            child: const Padding(
                              padding: EdgeInsets.all(8.0),
                              child: Icon(Icons.arrow_circle_up),
                            )),
                      ),
                    ),
                    Padding(
                      padding: const EdgeInsets.all(8.0),
                      child: RepeatOnTap(
                        call: () {
                          widget.service.send("tilt_down");
                        },
                        child: ElevatedButton(
                            onPressed: () {},
                            child: const Padding(
                              padding: EdgeInsets.all(8.0),
                              child: Icon(Icons.arrow_circle_down),
                            )),
                      ),
                    )
                  ],
                )
              ],
            ),
          ],
        ),
      ],
    );
  }
}
