# Ejemplos de envío/recepción para Android

## 📤 Crear Grupo (request)
```json
POST /grupos/
Content-Type: application/json

{
  "nombre": "Viaje Mendoza"
}
```

Response 201:
```json
{ "message": "Grupo registrado correctamente" }
```

Kotlin:
```kotlin
val json = """{"nombre": "Viaje Mendoza"}"""
val body = json.toRequestBody("application/json".toMediaType())
val req = Request.Builder().url("$baseUrl/grupos/").post(body).build()
val resp = client.newCall(req).execute()
```

## 📤 Crear Gasto
```json
POST /gastos/
Content-Type: application/json

{
  "descripcion": "Almuerzo",
  "monto": 25.50,
  "pagador_id": 1,
  "grupo_id": 1
}
```

Kotlin:
```kotlin
data class Gasto(val descripcion:String,val monto:Double,val pagador_id:Int,val grupo_id:Int)
val gasto = Gasto("Almuerzo",25.5,1,1)
val json = Gson().toJson(gasto)
val body = json.toRequestBody("application/json".toMediaType())
val req = Request.Builder().url("$baseUrl/gastos/").post(body).build()
client.newCall(req).execute()
```

## 📥 Recibir listado grupos
Kotlin:
```kotlin
val req = Request.Builder().url("$baseUrl/grupos/").get().build()
val resp = client.newCall(req).execute()
val arr = JSONArray(resp.body?.string() ?: "[]")
for(i in 0 until arr.length()){
  val obj = arr.getJSONObject(i)
  val id = obj.getInt("id")
  val nombre = obj.getString("nombre")
}
```

## 📥 WebSocket evento create
Servidor envía:
```json
{"type":"create","grupo_id":1,"data":{"descripcion":"Taxi","monto":12.5,"pagador_id":3,"grupo_id":1}}
```
Kotlin listener:
```kotlin
override fun onMessage(ws: WebSocket, text:String){
  val o=JSONObject(text)
  if(o.getString("type")=="create"){
    val d=o.getJSONObject("data")
    val desc=d.getString("descripcion")
    // etc
  }
}
```
