#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <ESP8266WebServer.h>
#include <ESP8266mDNS.h>
#include <Adafruit_Sensor.h>
#include <Adafruit_SSD1306.h>
#include <DHT.h>
#include <Wire.h>

#ifndef STASSID
#define STASSID ""
#define STAPSK  ""
#endif

#define SCREEN_WIDTH 128 // OLED display width, in pixels 
#define SCREEN_HEIGHT 64 // OLED display height, in pixels
#define SENSOR_READ_DELAY_MS 5000

const char* ssid = STASSID;
const char* password = STAPSK;

// Declaration for an SSD1306 display connected to I2C (SDA, SCL pins)
Adafruit_SSD1306 display(SCREEN_WIDTH, SCREEN_HEIGHT, &Wire, -1);

ESP8266WebServer server(80);

#define DHTPIN 5
#define DHTTYPE DHT11

DHT dht(DHTPIN, DHTTYPE);
float temperature = 0.0;
float humidity = 0.0;
float heatIndex = 0.0;

const int led = 13;

unsigned long previousSensorReadMillis = 0;

void handleRoot() {
  server.send(200, "text/plain", "hello from esp8266!");
}

void setup(void) {

  Serial.begin(115200);

  Wire.begin(2,14);
  if(!display.begin(SSD1306_SWITCHCAPVCC, 0x3c)) {
    Serial.println(F("SSD1306 allocation failed"));
    for(;;);
  }
  display.setTextColor(WHITE);

  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);
  Serial.println("");

  // Wait for connection
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("");
  Serial.print("Connected to ");
  Serial.println(ssid);
  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());

  if (MDNS.begin("esp8266")) {
    Serial.println("MDNS responder started");
  }

  server.on("/", handleRoot);

  server.on("/data", []() {
    char buffer[50];
    sprintf(buffer, "{\"temperature\": %.2f, \"heat_index\": %.2f, \"humidity\": %.2f}", temperature, heatIndex, humidity);
    server.send(200, "application/json", buffer);
  });

  server.on("/metrics", []() {
    char buffer[500];
    sprintf(buffer, "# HELP sensor_temperature_celsius The temperature in degrees censius.\n"
                    "# TYPE sensor_temperature_celsius gauge\n"
                    "sensor_temperature_celsius %.2f\n"
                    "# HELP sensor_humidity_percent The humidity percentage.\n"
                    "# TYPE sensor_humidity_percent gauge\n"
                    "sensor_humidity_percent %.2f\n"
                    "# HELP sensor_heat_index_celsius The humidity percentage.\n"
                    "# TYPE sensor_heat_index_celsius gauge\n"
                    "sensor_heat_index_celsius %.2f", temperature, humidity, heatIndex);

    server.send(200, "text/plain", buffer);
  });

  server.begin();
  Serial.println("HTTP server started");
}

void loop(void) {
  server.handleClient();
  MDNS.update();

  unsigned long currentTimeMillis = millis();
  if(currentTimeMillis - previousSensorReadMillis >= SENSOR_READ_DELAY_MS) {
    previousSensorReadMillis = currentTimeMillis;

    float newHumidity = dht.readHumidity();
    float newTemperature = dht.readTemperature();

    if (isnan(newHumidity) || isnan(newTemperature)) {
      Serial.println(F("Failed to read from  DHT sensor!"));
      return;
    }

    humidity = newHumidity;
    temperature = newTemperature;
    heatIndex = dht.computeHeatIndex(temperature, humidity, false);

    display.clearDisplay(); // display temperature
    display.setTextSize(1);
    display.setCursor(0,0);
    display.print("Temperature: ");
    display.setTextSize(2);
    display.setCursor(0,10);
    display.print(temperature);
    display.print(" ");
    display.setTextSize(1);
    display.cp437(true);
    display.write(167);
    display.setTextSize(2);
    display.print("C"); // display humidity
    display.setTextSize(1);
    display.setCursor(0, 35);
    display.print("Humidity: ");
    display.setTextSize(2);
    display.setCursor(0, 45);
    display.print(humidity);
    display.print(" %");
    display.display();
  }
}