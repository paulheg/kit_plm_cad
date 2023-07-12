import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

class Settings extends ChangeNotifier {
  String _host = "";
  get host => _host;

  Settings();

  Future<void> save() async {
    final prefs = await SharedPreferences.getInstance();

    await prefs.setString('host', _host);
  }

  void setHost(String host) {
    _host = host;
    notifyListeners();
  }

  Future<Settings> load() async {
    final prefs = await SharedPreferences.getInstance();

    _host = prefs.getString('host') ?? "";

    return this;
  }
}
