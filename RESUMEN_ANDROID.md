# 📱 Resumen Rápido - App Android Gastos

## 🔗 URLs Principales
```
API Base: http://44.197.255.1:8080

POST /grupos/      → Crear grupo
POST /gastos/      → Crear gasto
GET  /gastos/      → Listar gastos
ws   /ws           → WebSocket tiempo real
```

## ⚡ Orden de Pasos
```
1. Crear Grupo
   {"nombre": "Viaje"}
   
2. Crear Usuario
   {"user_name": "Juan", "email": "juan@email.com", "password": "123"}
   
3. Crear Gasto
   {"descripcion": "Almuerzo", "monto": 25.50, "pagador_id": 1, "grupo_id": 1}
   
4. Conectar WebSocket
   ws://44.197.255.1:8080/ws
   → Envía: {"type": "subscribe", "grupo_id": 1}
   → Recibe eventos: create, update, delete en tiempo real
```

## 📦 Dependencias Kotlin
```gradle
implementation 'com.squareup.okhttp3:okhttp:4.11.0'
implementation 'com.google.code.gson:gson:2.10.1'
implementation 'org.jetbrains.kotlinx:kotlinx-coroutines-android:1.7.3'
```

## 🔐 Permisos
```xml
<uses-permission android:name="android.permission.INTERNET" />
```

## 💾 2 Clases Necesarias

**ApiService.kt** → Llamadas REST (crear grupo, gasto, obtener datos)
**WebSocketManager.kt** → Conexión WebSocket (recibir cambios en tiempo real)

Ambas están completas en `GUIA_ANDROID.md` para copiar/pegar directamente.
