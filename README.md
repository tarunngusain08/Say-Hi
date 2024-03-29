# Say-Hi

## Minimum Viable Product (MVP) for Chat Application

### User Authentication
- User registration
- User login/logout
- Password hashing for security

### Messaging System
- Enable users to send text messages
- Real-time updates for new messages
- Message timestamp

### 1-1 Chat
- Create individual chat rooms for each pair of users
- Display sender and receiver information

### Error Handling
- Implement proper error handling for user inputs and server-side operations
- Provide meaningful error messages to users


## Database Design

For our chat application, we'll employ a hybrid approach, utilizing both traditional relational databases and blob storage for efficient data management.

1. **Relational Database:**
   - **Entities:**
     - **User:** Stores user information, including user details and profile settings.
     - **Message:** Manages message metadata such as sender, receiver, timestamp, and references to message content.

2. **Blob Storage:**
   - Stores large message bodies, file attachments, and other binary data.
   - Efficiently handles scalability and performance for non-relational data.


## Entity-Relationship (ER) Relationships:**

1. **User to Message (1:N):**
   - One user can send multiple messages.

2. **User to Chat History (N:N):**
   - Many users can have many chat histories, each associated with a specific chat partner.
