import 'dart:convert';

import 'package:kerberos_control/api/connection_status.dart';
import 'package:http/http.dart' as http;

Future<List<ConnectionStatus>> getConnectionStatusList(String host) async {
  var endpoint = Uri.tryParse("https://${host}/robots");
  if (endpoint == null) {
    return Future.error(Exception("Could not parse endpoint"));
  }

  try {
    var response = await http.get(endpoint);
    Iterable list = jsonDecode(response.body);
    List<ConnectionStatus> cons = List<ConnectionStatus>.from(
        list.map((e) => ConnectionStatus.fromJson(e)));
    return cons;
  } catch (e) {
    rethrow;
  }
}
