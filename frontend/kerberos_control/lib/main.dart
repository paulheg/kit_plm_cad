import 'package:kerberos_control/pages/connect_page.dart';
import 'package:kerberos_control/pages/control_page.dart';
import 'package:kerberos_control/pages/select_robot_page.dart';
import 'package:kerberos_control/settings.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:go_router/go_router.dart';

void main() {
  runApp(MultiProvider(
    providers: [
      ChangeNotifierProvider<Settings>(
        create: (context) => Settings(),
      )
    ],
    child: const MyApp(),
  ));
}

final _router = GoRouter(
  initialLocation: '/connect',
  routes: [
    GoRoute(
      path: '/connect',
      builder: (context, state) => const ConnectPage(),
    ),
    GoRoute(
      name: 'control',
      path: '/control/:id',
      builder: (context, state) {
        String robotId = state.pathParameters['id']!;

        return ControlPage(robotId);
      },
    ),
    GoRoute(
      path: '/select',
      builder: (context, state) => const SelectRobotPage(),
    )
  ],
);

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      routerConfig: _router,
    );
  }
}
