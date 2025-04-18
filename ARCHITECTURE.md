# OverHelloWorld Architecture

This document describes the (over)engineered architecture of the OverHelloWorld project.

## Overview Diagram

```
+-------------------+      +---------------------+      +-------------+
|    API Layer      +----->|  Command Handler    +----->|  Event Bus  |
+-------------------+      +---------------------+      +------+------+ 
                                                         |      |      
                                                         v      v      
                                                      +----+  +----+  +---------+
                                                      |Redis|  |File|  |Plugins |
                                                      +----+  +----+  +---------+
                                                        |        |        |
                                                        v        v        v
                                                    +-------------------------+
                                                    |    Read Model           |
                                                    +-------------------------+
                                                         |
                                                         v
                                                  +-----------------+
                                                  |  Observability  |
                                                  +-----------------+
```

## Components
- **API Layer**: Handles HTTP requests for commands and queries.
- **Command Handler**: Implements CQRS, dispatches commands, and triggers events.
- **Event Bus**: Publishes events to Redis, persists to file, and notifies plugins.
- **Redis**: Used for event pub/sub.
- **File Event Store**: Persists all events for replayability.
- **Plugins**: ASCII art, TTS, LED, and more for maximum overengineering.
- **Read Model**: Query-optimized view of events.
- **Observability**: Logging, metrics, and tracing everywhere.

## Notes
- Designed for maximum complexity, minimum practicality.
- Every component is extensible, even if unnecessary.

---

For more details, see inline comments and [OVERENGINEERING_ARTICLE.md](./OVERENGINEERING_ARTICLE.md).
