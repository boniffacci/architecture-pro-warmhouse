from flask import Flask, request, jsonify

app = Flask(__name__)
devices = {}

@app.route('/devices/<device_id>', methods=['POST'])
def create_device(device_id):
    devices[device_id] = {"status": "off"}
    return jsonify({"message": f"Device {device_id} created"}), 201

@app.route('/devices/<device_id>/status', methods=['PATCH'])
def update_status(device_id):
    data = request.json
    if device_id not in devices:
        return jsonify({"error": "Device not found"}), 404
    devices[device_id]["status"] = data.get("status", "off")
    return jsonify({"device_id": device_id, "status": devices[device_id]["status"]})

@app.route('/devices/<device_id>', methods=['GET'])
def get_device(device_id):
    device = devices.get(device_id)
    if not device:
        return jsonify({"error": "Device not found"}), 404
    return jsonify({"device_id": device_id, "status": device["status"]})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5002)
