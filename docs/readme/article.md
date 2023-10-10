# REST API Article

## list articles

### GET ```127.0.0.1:4030/articles```
### Response
```json
{
    "results": [
        {
            "id": "6",
            "question_id": "6",
            "created_at": "2023-05-12T11:15:17.413858Z",
            "title": "How Blockchain Can Prevent Device Tampering",
            "content": "In today's world, where technology is advancing at an unprecedented pace, the security of devices has become a major concern. With the rise of the Internet of Things (IoT), the number of connected devices has increased significantly, making them vulnerable to cyber-attacks. One of the most significant threats to device security is tampering. However, blockchain technology can provide a solution to this problem.\n\nDevice tampering is a process where an attacker gains unauthorized access to a device and modifies its functionality or data. This can be done by physically altering the device or by exploiting vulnerabilities in its software.",
            "status": "publish"
        },
      {
        "id": "6",
        "question_id": "6",
        "created_at": "2023-05-12T11:15:17.413858Z",
        "title": "How Blockchain Can Prevent Device Tampering",
        "content": "In today's world, where technology is advancing at an unprecedented pace, the security of devices has become a major concern. With the rise of the Internet of Things (IoT), the number of connected devices has increased significantly, making them vulnerable to cyber-attacks. One of the most significant threats to device security is tampering. However, blockchain technology can provide a solution to this problem.\n\nDevice tampering is a process where an attacker gains unauthorized access to a device and modifies its functionality or data. This can be done by physically altering the device or by exploiting vulnerabilities in its software.",
        "status": "draft"
      }
    ]
}
```

### GET ```127.0.0.1:4030/articles?status=publish```
### Response
```json
{
    "results": [
        {
            "id": "6",
            "question_id": "6",
            "created_at": "2023-05-12T11:15:17.413858Z",
            "title": "How Blockchain Can Prevent Device Tampering",
            "content": "In today's world, where technology is advancing at an unprecedented pace, the security of devices has become a major concern. With the rise of the Internet of Things (IoT), the number of connected devices has increased significantly, making them vulnerable to cyber-attacks. One of the most significant threats to device security is tampering. However, blockchain technology can provide a solution to this problem.\n\nDevice tampering is a process where an attacker gains unauthorized access to a device and modifies its functionality or data. This can be done by physically altering the device or by exploiting vulnerabilities in its software.",
            "status": "publish"
        }
    ]
}
```

### GET ```127.0.0.1:4030/articles?status=draft```
### Response
```json
{
    "results": [
        {
            "id": "6",
            "question_id": "6",
            "created_at": "2023-05-12T11:15:17.413858Z",
            "title": "How Blockchain Can Prevent Device Tampering",
            "content": "In today's world, where technology is advancing at an unprecedented pace, the security of devices has become a major concern. With the rise of the Internet of Things (IoT), the number of connected devices has increased significantly, making them vulnerable to cyber-attacks. One of the most significant threats to device security is tampering. However, blockchain technology can provide a solution to this problem.\n\nDevice tampering is a process where an attacker gains unauthorized access to a device and modifies its functionality or data. This can be done by physically altering the device or by exploiting vulnerabilities in its software.",
            "status": "draft"
        }
    ]
}
```
