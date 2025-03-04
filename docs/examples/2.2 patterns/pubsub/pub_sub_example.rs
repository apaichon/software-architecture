use std::collections::HashMap;
use std::sync::{Arc, Mutex};

// Define message type for type-safety
#[derive(Clone, Debug)]
struct Message {
    topic: String,
    content: String,
}

// Publisher trait defining behavior for publishing messages
trait Publisher {
    fn publish(&self, topic: &str, content: &str);
}

// Subscriber trait defining behavior for handling messages
trait Subscriber {
    fn handle_message(&self, message: &Message);
}

// Message broker to manage pub/sub relationships
struct MessageBroker {
    subscribers: Arc<Mutex<HashMap<String, Vec<Box<dyn Subscriber + Send>>>>>,
}

impl MessageBroker {
    fn new() -> Self {
        MessageBroker {
            subscribers: Arc::new(Mutex::new(HashMap::new())),
        }
    }

    // Subscribe to a topic
    fn subscribe(&self, topic: &str, subscriber: Box<dyn Subscriber + Send>) {
        let mut subs = self.subscribers.lock().unwrap();
        subs.entry(topic.to_string())
            .or_insert_with(Vec::new)
            .push(subscriber);
    }

    // Implement Publisher trait for MessageBroker
    impl Publisher for MessageBroker {
        fn publish(&self, topic: &str, content: &str) {
            let message = Message {
                topic: topic.to_string(),
                content: content.to_string(),
            };
            
            let subs = self.subscribers.lock().unwrap();
            if let Some(topic_subscribers) = subs.get(topic) {
                for subscriber in topic_subscribers {
                    subscriber.handle_message(&message);
                }
            }
        }
    }
}

// Example subscriber implementation
struct LogSubscriber {
    name: String,
}

impl LogSubscriber {
    fn new(name: &str) -> Self {
        LogSubscriber {
            name: name.to_string(),
        }
    }
}

impl Subscriber for LogSubscriber {
    fn handle_message(&self, message: &Message) {
        println!(
            "Subscriber {} received message on topic '{}': {}",
            self.name, message.topic, message.content
        );
    }
}

// Example usage
fn main() {
    // Create message broker
    let broker = MessageBroker::new();

    // Create subscribers
    let subscriber1 = Box::new(LogSubscriber::new("Logger1"));
    let subscriber2 = Box::new(LogSubscriber::new("Logger2"));

    // Subscribe to topics
    broker.subscribe("news", subscriber1);
    broker.subscribe("news", subscriber2);

    // Publish messages
    broker.publish("news", "Breaking: New Rust Version Released!");

    // Output will show both subscribers receiving the message
}

// Benefits of this pattern:
// 1. Decoupling: Publishers and subscribers are loosely coupled
// 2. Scalability: Easy to add new subscribers without modifying publishers
// 3. Flexibility: Subscribers can handle messages differently
// 4. Threading: Can be extended for async/concurrent processing
//
// Use cases:
// 1. Event-driven architectures
// 2. Distributed systems
// 3. Message queuing systems
// 4. Real-time data processing
