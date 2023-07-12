import 'package:web_socket_channel/web_socket_channel.dart';

class ConnectionService {
  WebSocketChannel? _channel;

  Future<ConnectionService> connect(
      String host, String robotId, void Function()? onDone) async {
    if (_channel != null) return this;

    final channel = WebSocketChannel.connect(
      Uri.parse('wss://$host/ws/remote/$robotId'),
    );

    await channel.ready;

    channel.stream.listen((event) {}, onDone: onDone);

    _channel = channel;
    return this;
  }

  void send(String message) {
    if (_channel != null && _channel!.closeCode == null) {
      _channel!.sink.add(message);
    }
  }

  void disconnect() {
    if (_channel != null) {
      _channel!.sink.close();
    }
  }
}
