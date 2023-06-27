 #include <Servo.h>

#include <AntaresESP32HTTP.h>


#define ACCESSKEY "2c4b04aa13dbc9bc:cb8f6d670f09d6fc"
#define WIFISSID "Gru"
#define PASSWORD "girama25"

#define projectName "smartTrashCan"
#define deviceName "sensor-dalam"

AntaresESP32HTTP antares(ACCESSKEY);

#define ECHO_PIN_LUAR 14
#define TRIG_PIN_LUAR 12
#define TRIG_PIN_DALAM 18
#define ECHO_PIN_DALAM 5
#define LED_HIJAU 32
#define LED_MERAH 33



Servo servo1;

// the setup function runs once when you press reset or power the board
void setup() {
  Serial.begin(115200);  
  pinMode(LED_HIJAU, OUTPUT);
  pinMode(LED_MERAH, OUTPUT);
  pinMode(TRIG_PIN_LUAR, OUTPUT);
  pinMode(ECHO_PIN_LUAR, INPUT);
  pinMode(ECHO_PIN_DALAM, INPUT);
  pinMode(TRIG_PIN_DALAM, OUTPUT);
  servo1.attach(13);
  antares.setDebug(true);
  antares.wifiConnection(WIFISSID, PASSWORD);
}

int readDistanceCMLuar() {
  digitalWrite(TRIG_PIN_LUAR, LOW);
  delayMicroseconds(2);
  digitalWrite(TRIG_PIN_LUAR, HIGH);
  delayMicroseconds(10);
  digitalWrite(TRIG_PIN_LUAR, LOW);
  int duration = pulseIn(ECHO_PIN_LUAR, HIGH);
  return duration * 0.034 / 2;
}
int readDistanceCMDalam() {
  digitalWrite(TRIG_PIN_DALAM, LOW);
  delayMicroseconds(2);
  digitalWrite(TRIG_PIN_DALAM, HIGH);
  delayMicroseconds(10);
  digitalWrite(TRIG_PIN_DALAM, LOW);
  int duration = pulseIn(ECHO_PIN_DALAM, HIGH);
  return duration * 0.034 / 2;
  
}

// the loop function runs over and over again forever
void loop() {
  float position = 0.0;
  float distanceLuar = readDistanceCMLuar();
  bool isNearbyLuar = distanceLuar <= 50;
  float distanceDalam = readDistanceCMDalam();
  bool isNearbyDalam = distanceDalam <=5;
  int kapasitas = distanceDalam*4-4;
  String statusKapasitas;
  long durasiSerial = 0;
  long durasiSensor = 0;
  long jedaSerial = 500;
  long jedaSensor = 10;
  
  if(millis() - durasiSensor > jedaSensor){
    if(isNearbyDalam ){
      digitalWrite(LED_MERAH, HIGH);
      digitalWrite(LED_HIJAU, LOW);
    }
    else{
      digitalWrite(LED_HIJAU, HIGH);
      digitalWrite(LED_MERAH, LOW);
    }
  
    if(isNearbyLuar && isNearbyDalam){
      position = 0;
      servo1.write(position);
    }
    else if (isNearbyLuar){
      position += 180;
      servo1.write(position);
      delay(4000);
    }
    else{
      position = 0;
      servo1.write(position);
    }
  
    
    Serial.print("Measured distance Luar: ");
    Serial.println(readDistanceCMLuar());
    Serial.print("Kapasitas Tempat Sampah: ");
    Serial.println(kapasitas);
  
    if (kapasitas<17){
      statusKapasitas = "Penuh";
      Serial.print(statusKapasitas);
  
    }
    else{
      statusKapasitas = "Tersedia";
      Serial.print(statusKapasitas);
      
    } 
    durasiSensor = millis();
  }
  
  if(millis() - durasiSerial > jedaSerial){
      antares.add("kapasitas", kapasitas);
      antares.add("status", statusKapasitas);
      antares.send(projectName, deviceName);
      durasiSerial = millis();
    }
}