import {
  WifiCredentials,
  Confirmation,
  Request
} from "./pb/wifi_service_pb.js";
import { WifiServiceClient } from "./pb/wifi_service_grpc_web_pb.js";

const client = new WifiServiceClient("http://localhost:8080");

function sendCredentialsCallback(err, response) {
  if (err) {
    console.error(err.message);
    return;
  }
  const confirmation = response.toObject();
  console.log("Sucessfull password: ", confirmation);

  // Return message
  document.querySelector("#confirmation").innerHTML = confirmation.message
}

function saveCredentials() {
  // Build request
  const ssid = document.querySelector("#ssid").value;
  const password = document.querySelector("#password").value;
  if (!ssid || !password) return;
  const request = new WifiCredentials();
  request.setSsid(ssid);
  request.setPassword(password);

  // Clear password
  document.querySelector("#password").value = "";

  // Make request
  client.sendCredentials(request, {}, sendCredentialsCallback);
}

// Make function globally available
window.saveCredentials = saveCredentials;
