import 'dart:io';

import 'package:flutter/material.dart';

class RepeatOnTap extends StatefulWidget {
  final int? milliseconds;
  final VoidCallback? call;
  final Widget? child;

  const RepeatOnTap({this.child, this.call, this.milliseconds, super.key});

  @override
  State<RepeatOnTap> createState() => _RepeatOnTapState();
}

class _RepeatOnTapState extends State<RepeatOnTap> {
  bool repeat = false;

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTapDown: (details) async {
        if (repeat) {
          return;
        }

        repeat = true;

        while (repeat) {
          if (widget.call != null) {
            widget.call!.call();
          }

          await Future.delayed(
              Duration(milliseconds: widget.milliseconds ?? 100));
        }
      },
      onTapUp: (details) {
        repeat = false;
      },
      onTapCancel: () {
        repeat = false;
      },
      child: widget.child,
    );
  }
}
