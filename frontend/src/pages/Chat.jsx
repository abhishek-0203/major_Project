import React, { useState, useEffect, useRef } from 'react';
import {
  Box,
  TextField,
  Button,
  Paper,
  Typography,
  List,
  ListItem,
  ListItemText,
  Container,
  Divider,
} from '@mui/material';
import { Send } from '@mui/icons-material';

const Chat = () => {
  const [messages, setMessages] = useState([]);
  const [newMessage, setNewMessage] = useState('');
  const [ws, setWs] = useState(null);
  const messagesEndRef = useRef(null);

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  };

  useEffect(() => {
    const token = localStorage.getItem('token');
    const socket = new WebSocket(`ws://localhost:8081/ws/chat?token=${token}`);

    socket.onopen = () => {
      console.log('Connected to chat');
    };

    socket.onmessage = (event) => {
      const message = JSON.parse(event.data);
      setMessages((prevMessages) => [...prevMessages, message]);
    };

    socket.onclose = () => {
      console.log('Disconnected from chat');
    };

    setWs(socket);

    return () => {
      socket.close();
    };
  }, []);

  useEffect(() => {
    scrollToBottom();
  }, [messages]);

  const handleSendMessage = (e) => {
    e.preventDefault();
    if (newMessage.trim() && ws) {
      const message = {
        content: newMessage,
        timestamp: new Date().toISOString(),
        sender: localStorage.getItem('userRole'),
      };
      ws.send(JSON.stringify(message));
      setNewMessage('');
    }
  };

  return (
    <Container maxWidth="md">
      <Paper elevation={3} sx={{ mt: 4, height: '80vh', display: 'flex', flexDirection: 'column' }}>
        <Box sx={{ p: 2, bgcolor: 'primary.main', color: 'white' }}>
          <Typography variant="h6">Chat Room</Typography>
        </Box>
        <Box sx={{ flexGrow: 1, overflow: 'auto', p: 2 }}>
          <List>
            {messages.map((message, index) => (
              <ListItem
                key={index}
                sx={{
                  display: 'flex',
                  justifyContent: message.sender === localStorage.getItem('userRole')
                    ? 'flex-end'
                    : 'flex-start',
                }}
              >
                <Paper
                  sx={{
                    p: 2,
                    bgcolor: message.sender === localStorage.getItem('userRole')
                      ? 'primary.light'
                      : 'grey.100',
                    maxWidth: '70%',
                  }}
                >
                  <ListItemText
                    primary={message.content}
                    secondary={new Date(message.timestamp).toLocaleString()}
                  />
                </Paper>
              </ListItem>
            ))}
            <div ref={messagesEndRef} />
          </List>
        </Box>
        <Divider />
        <Box sx={{ p: 2, bgcolor: 'background.default' }}>
          <form onSubmit={handleSendMessage} style={{ display: 'flex', gap: 8 }}>
            <TextField
              fullWidth
              value={newMessage}
              onChange={(e) => setNewMessage(e.target.value)}
              placeholder="Type a message"
              variant="outlined"
              size="small"
            />
            <Button
              type="submit"
              variant="contained"
              endIcon={<Send />}
              disabled={!newMessage.trim()}
            >
              Send
            </Button>
          </form>
        </Box>
      </Paper>
    </Container>
  );
};

export default Chat;