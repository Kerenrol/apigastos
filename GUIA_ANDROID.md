# 📱 Guía Completa - App Android para Gastos en Tiempo Real

## 🔗 URLs Base de tu API

**Servidor:** `http://44.197.255.1:8080`

### Endpoints que necesitas:

| Función | Método | URL | 
|---------|--------|-----|
| **Crear Grupo** | POST | `http://44.197.255.1:8080/grupos/` |
| **Listar Grupos** | GET | `http://44.197.255.1:8080/grupos/` |
| **Crear Gasto** | POST | `http://44.197.255.1:8080/gastos/` |
| **Listar Gastos** | GET | `http://44.197.255.1:8080/gastos/` |
| **WebSocket** | WS | `ws://44.197.255.1:8080/ws` |

---

## 🎯 Flujo Correcto (Paso a Paso)

```
1. CREAR GRUPO (requerido primero)
   ↓
2. CREAR USUARIO/PAGADOR (si aún no existe)
   ↓
3. CREAR GASTO
   ↓
4. CONECTAR A WEBSOCKET (para actualizaciones en tiempo real)
```

---

## 📋 Paso 1: Crear Grupo

**Endpoint:**
```
POST http://44.197.255.1:8080/grupos/
Content-Type: application/json

{
  "nombre": "Viaje Mendoza"
}
```

**Respuesta (éxito 201):**
```json
{
  "message": "Grupo registrado correctamente"
}
```

**En Android (Kotlin):**
```kotlin
suspend fun crearGrupo(nombreGrupo: String): Int? {
    return try {
        val client = OkHttpClient()
        val requestBody = """{"nombre": "$nombreGrupo"}""".toRequestBody("application/json".toMediaType())
        
        val request = Request.Builder()
            .url("http://44.197.255.1:8080/grupos/")
            .post(requestBody)
            .build()
        
        val response = client.newCall(request).execute()
        if (response.isSuccessful) {
            // El grupo se creó, obtén su ID consultando la lista
            obtenerUltimoGrupoId()
        } else null
    } catch (e: Exception) {
        Log.e("CrearGrupo", e.message ?: "Error desconocido")
        null
    }
}

suspend fun obtenerUltimoGrupoId(): Int? {
    return try {
        val client = OkHttpClient()
        val request = Request.Builder()
            .url("http://44.197.255.1:8080/grupos/")
            .get()
            .build()
        
        val response = client.newCall(request).execute()
        val jsonArray = JSONArray(response.body?.string() ?: "[]")
        
        if (jsonArray.length() > 0) {
            jsonArray.getJSONObject(jsonArray.length() - 1).getInt("id")
        } else null
    } catch (e: Exception) {
        Log.e("ObtenerGrupo", e.message ?: "Error desconocido")
        null
    }
}
```

---

## 👤 Paso 2: Crear Usuario (Pagador)

**Endpoint:**
```
POST http://44.197.255.1:8080/users/
Content-Type: application/json

{
  "user_name": "Juan Pérez",
  "email": "juan@example.com",
  "password": "miPassword123"
}
```

**Respuesta (éxito 201):**
```json
{
  "message": "Usuario registrado correctamente"
}
```

**En Android (Kotlin):**
```kotlin
suspend fun crearUsuario(nombre: String, email: String, contraseña: String): Int? {
    return try {
        val client = OkHttpClient()
        val json = """
            {
              "user_name": "$nombre",
              "email": "$email",
              "password": "$contraseña"
            }
        """.trimIndent()
        
        val requestBody = json.toRequestBody("application/json".toMediaType())
        val request = Request.Builder()
            .url("http://44.197.255.1:8080/users/")
            .post(requestBody)
            .build()
        
        val response = client.newCall(request).execute()
        if (response.isSuccessful) {
            obtenerUltimoUsuarioId()
        } else null
    } catch (e: Exception) {
        Log.e("CrearUsuario", e.message ?: "Error desconocido")
        null
    }
}
```

---

## 💰 Paso 3: Crear Gasto

**Endpoint:**
```
POST http://44.197.255.1:8080/gastos/
Content-Type: application/json

{
  "descripcion": "Almuerzo",
  "monto": 25.50,
  "pagador_id": 1,
  "grupo_id": 1
}
```

**Respuesta (éxito 201):**
```json
{
  "message": "Gasto registrado correctamente"
}
```

**En Android (Kotlin):**
```kotlin
data class Gasto(
    val descripcion: String,
    val monto: Double,
    val pagador_id: Int,
    val grupo_id: Int
)

suspend fun crearGasto(gasto: Gasto): Boolean {
    return try {
        val client = OkHttpClient()
        val json = Gson().toJson(gasto)
        val requestBody = json.toRequestBody("application/json".toMediaType())
        
        val request = Request.Builder()
            .url("http://44.197.255.1:8080/gastos/")
            .post(requestBody)
            .build()
        
        val response = client.newCall(request).execute()
        response.isSuccessful
    } catch (e: Exception) {
        Log.e("CrearGasto", e.message ?: "Error desconocido")
        false
    }
}
```

---

## 🔄 Paso 4: Conectar a WebSocket (Tiempo Real)

**URL WebSocket:**
```
ws://44.197.255.1:8080/ws
```

**En Android (Kotlin) con OkHttp WebSocket:**
```kotlin
import okhttp3.*
import okhttp3.WebSocket
import okhttp3.WebSocketListener

class GastoWebSocketListener(private val onEvent: (evento: String) -> Unit) : WebSocketListener() {
    
    override fun onOpen(webSocket: WebSocket, response: Response) {
        Log.d("WebSocket", "Conectado")
        // Suscribirse a un grupo
        val mensaje = """{"type": "subscribe", "grupo_id": 1}"""
        webSocket.send(mensaje)
    }
    
    override fun onMessage(webSocket: WebSocket, text: String) {
        Log.d("WebSocket", "Mensaje: $text")
        onEvent(text)
        // Actualiza tu UI aquí con los cambios en tiempo real
    }
    
    override fun onFailure(webSocket: WebSocket, t: Throwable, response: Response?) {
        Log.e("WebSocket", "Error: ${t.message}")
    }
    
    override fun onClosed(webSocket: WebSocket, code: Int, reason: String) {
        Log.d("WebSocket", "Desconectado: $reason")
    }
}

// Uso:
fun conectarWebSocket(grupoId: Int) {
    val client = OkHttpClient()
    val request = Request.Builder()
        .url("ws://44.197.255.1:8080/ws")
        .build()
    
    val webSocket = client.newWebSocket(
        request,
        GastoWebSocketListener { evento ->
            // Procesa los eventos aquí
            procesarEvento(evento)
        }
    )
}

fun procesarEvento(eventoJson: String) {
    try {
        val jsonObject = JSONObject(eventoJson)
        val tipo = jsonObject.getString("type") // "create", "update", "delete"
        val datos = jsonObject.getJSONObject("data")
        
        when (tipo) {
            "create" -> {
                Log.d("Evento", "Nuevo gasto creado")
                // Actualiza tu lista de gastos
            }
            "update" -> {
                Log.d("Evento", "Gasto actualizado")
                // Actualiza el gasto en tu lista
            }
            "delete" -> {
                Log.d("Evento", "Gasto eliminado")
                // Elimina el gasto de tu lista
            }
        }
    } catch (e: Exception) {
        Log.e("Evento", e.message ?: "Error procesando evento")
    }
}
```

---

## 🏗️ Estructura Recomendada del Proyecto Android

```
MyGastosApp/
├── app/
│   ├── src/
│   │   ├── main/
│   │   │   ├── java/com/example/miapp/
│   │   │   │   ├── api/
│   │   │   │   │   ├── ApiService.kt        # Llamadas REST
│   │   │   │   │   └── WebSocketManager.kt  # Manejo WebSocket
│   │   │   │   ├── models/
│   │   │   │   │   ├── Grupo.kt
│   │   │   │   │   ├── Gasto.kt
│   │   │   │   │   └── Usuario.kt
│   │   │   │   ├── ui/
│   │   │   │   │   ├── GastosActivity.kt    # Pantalla principal
│   │   │   │   │   ├── NuevoGastoActivity.kt
│   │   │   │   │   └── NuevoGrupoActivity.kt
│   │   │   │   └── MainActivity.kt
│   │   │   └── AndroidManifest.xml
│   │   └── res/
│   └── build.gradle
└── settings.gradle
```

---

## 📦 Dependencias (build.gradle)

```gradle
dependencies {
    // OkHttp para REST y WebSocket
    implementation 'com.squareup.okhttp3:okhttp:4.11.0'
    
    // Gson para JSON
    implementation 'com.google.code.gson:gson:2.10.1'
    
    // Coroutines
    implementation 'org.jetbrains.kotlinx:kotlinx-coroutines-android:1.7.3'
    implementation 'org.jetbrains.kotlinx:kotlinx-coroutines-core:1.7.3'
    
    // Lifecycle
    implementation 'androidx.lifecycle:lifecycle-runtime-ktx:2.6.2'
    
    // Retrofit (alternativa a OkHttp)
    implementation 'com.squareup.retrofit2:retrofit:2.10.0'
    implementation 'com.squareup.retrofit2:converter-gson:2.10.0'
}
```

---

## 🔐 Permisos (AndroidManifest.xml)

```xml
<uses-permission android:name="android.permission.INTERNET" />
<uses-permission android:name="android.permission.ACCESS_NETWORK_STATE" />
```

---

## 📱 Clase Completa - ApiService.kt

```kotlin
import android.util.Log
import com.google.gson.Gson
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.RequestBody.Companion.toRequestBody
import org.json.JSONArray
import org.json.JSONObject

class ApiService {
    private val baseUrl = "http://44.197.255.1:8080"
    private val client = OkHttpClient()     
    private val gson = Gson()

    // GRUPOS
    suspend fun crearGrupo(nombre: String): Boolean = try {
        val json = """{"nombre": "$nombre"}"""
        val requestBody = json.toRequestBody("application/json".toMediaType())
        val request = Request.Builder()
            .url("$baseUrl/grupos/")
            .post(requestBody)
            .build()
        
        client.newCall(request).execute().isSuccessful
    } catch (e: Exception) {
        Log.e("ApiService", "Error creando grupo: ${e.message}")
        false
    }

    suspend fun obtenerGrupos(): List<Map<String, Any>>? = try {
        val request = Request.Builder()
            .url("$baseUrl/grupos/")
            .get()
            .build()
        
        val response = client.newCall(request).execute()
        if (response.isSuccessful) {
            val jsonArray = JSONArray(response.body?.string() ?: "[]")
            val grupos = mutableListOf<Map<String, Any>>()
            
            for (i in 0 until jsonArray.length()) {
                val obj = jsonArray.getJSONObject(i)
                grupos.add(mapOf(
                    "id" to obj.getInt("id"),
                    "nombre" to obj.getString("nombre")
                ))
            }
            grupos
        } else null
    } catch (e: Exception) {
        Log.e("ApiService", "Error obteniendo grupos: ${e.message}")
        null
    }

    // GASTOS
    suspend fun crearGasto(
        descripcion: String,
        monto: Double,
        pagadorId: Int,
        grupoId: Int
    ): Boolean = try {
        val json = """
            {
              "descripcion": "$descripcion",
              "monto": $monto,
              "pagador_id": $pagadorId,
              "grupo_id": $grupoId
            }
        """.trimIndent()
        
        val requestBody = json.toRequestBody("application/json".toMediaType())
        val request = Request.Builder()
            .url("$baseUrl/gastos/")
            .post(requestBody)
            .build()
        
        client.newCall(request).execute().isSuccessful
    } catch (e: Exception) {
        Log.e("ApiService", "Error creando gasto: ${e.message}")
        false
    }

    suspend fun obtenerGastos(): List<Map<String, Any>>? = try {
        val request = Request.Builder()
            .url("$baseUrl/gastos/")
            .get()
            .build()
        
        val response = client.newCall(request).execute()
        if (response.isSuccessful) {
            val jsonArray = JSONArray(response.body?.string() ?: "[]")
            val gastos = mutableListOf<Map<String, Any>>()
            
            for (i in 0 until jsonArray.length()) {
                val obj = jsonArray.getJSONObject(i)
                gastos.add(mapOf(
                    "id" to obj.getInt("id"),
                    "descripcion" to obj.getString("descripcion"),
                    "monto" to obj.getDouble("monto"),
                    "pagador_id" to obj.getInt("pagador_id"),
                    "grupo_id" to obj.getInt("grupo_id"),
                    "fecha" to obj.getString("fecha")
                ))
            }
            gastos
        } else null
    } catch (e: Exception) {
        Log.e("ApiService", "Error obteniendo gastos: ${e.message}")
        null
    }
}
```

---

## 🎬 Clase Completa - WebSocketManager.kt

```kotlin
import android.util.Log
import okhttp3.*
import org.json.JSONObject

class WebSocketManager(
    private val onGastoCreado: (datos: Map<String, Any>) -> Unit,
    private val onGastoActualizado: (datos: Map<String, Any>) -> Unit,
    private val onGastoEliminado: (id: Int) -> Unit
) {
    private val client = OkHttpClient()
    private var webSocket: WebSocket? = null
    private val baseUrl = "ws://44.197.255.1:8080/ws"

    fun conectar(grupoId: Int) {
        val request = Request.Builder()
            .url(baseUrl)
            .build()
        
        webSocket = client.newWebSocket(request, object : WebSocketListener() {
            override fun onOpen(webSocket: WebSocket, response: Response) {
                Log.d("WebSocket", "Conectado")
                suscribirse(grupoId)
            }

            override fun onMessage(webSocket: WebSocket, text: String) {
                procesarMensaje(text)
            }

            override fun onFailure(webSocket: WebSocket, t: Throwable, response: Response?) {
                Log.e("WebSocket", "Error: ${t.message}")
            }

            override fun onClosed(webSocket: WebSocket, code: Int, reason: String) {
                Log.d("WebSocket", "Desconectado")
            }
        })
    }

    private fun suscribirse(grupoId: Int) {
        val mensaje = """{"type": "subscribe", "grupo_id": $grupoId}"""
        webSocket?.send(mensaje)
    }

    private fun procesarMensaje(texto: String) {
        try {
            val json = JSONObject(texto)
            val tipo = json.getString("type")
            val datos = json.getJSONObject("data")
            
            when (tipo) {
                "create" -> {
                    val datosMapa = mapOf(
                        "descripcion" to datos.getString("descripcion"),
                        "monto" to datos.getDouble("monto"),
                        "pagador_id" to datos.getInt("pagador_id"),
                        "grupo_id" to datos.getInt("grupo_id")
                    )
                    onGastoCreado(datosMapa)
                }
                "update" -> {
                    val datosMapa = mapOf(
                        "id" to datos.getInt("id"),
                        "descripcion" to datos.getString("descripcion"),
                        "monto" to datos.getDouble("monto")
                    )
                    onGastoActualizado(datosMapa)
                }
                "delete" -> {
                    val id = datos.getInt("id")
                    onGastoEliminado(id)
                }
            }
        } catch (e: Exception) {
            Log.e("WebSocket", "Error procesando mensaje: ${e.message}")
        }
    }

    fun desconectar() {
        webSocket?.close(1000, "Cierre normal")
    }
}
```

---

## 📋 Resumen de URLs

```
Base URL: http://44.197.255.1:8080

REST:
- POST   /grupos/          → Crear grupo
- GET    /grupos/          → Obtener todos los grupos
- POST   /gastos/          → Crear gasto
- GET    /gastos/          → Obtener todos los gastos

WebSocket:
- ws://44.197.255.1:8080/ws → Conexión en tiempo real

Ejemplo:
POST http://44.197.255.1:8080/gastos/
{
  "descripcion": "Almuerzo",
  "monto": 25.50,
  "pagador_id": 1,
  "grupo_id": 1
}
```

---

¿Necesitas ayuda con algo específico de Android?
