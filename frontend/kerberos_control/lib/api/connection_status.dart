class ConnectionStatus {
  String id;
  bool connected;

  ConnectionStatus(this.id, this.connected);

  ConnectionStatus.fromJson(Map<String, dynamic> json)
      : id = json['robot_id'],
        connected = json['connected'];
}
