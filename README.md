# Cache-Store  

`cache-store` is a lightweight, high-performance **in-memory key-value store** that works similarly to **Redis**.  
It allows you to **store**, **retrieve**, and **persist** data efficiently, making it suitable for caching, session management, and fast data lookups.  

---

## 🚀 Features  

- ⚡ **In-Memory Storage** – Store and retrieve data at lightning speed.  
- 🔑 **Key-Value Access** – Simple API for `SET` and `GET` operations.  
- 💾 **Persistence** – Save in-memory data to disk as a **data dump** and reload it later.  
- 🔄 **Redis-like Behavior** – Familiar usage for those who have worked with Redis.  
- 🛠️ **Lightweight & Easy-to-Use** – Minimal setup and easy integration.  

---

## ⚙️ How It Works  

1. **Store Data in Memory**  
   Data is stored in memory for ultra-fast access.  

   ```go
   cache.Set("user:1", "Deepali")
