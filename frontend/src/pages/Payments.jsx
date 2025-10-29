import React, { useEffect, useState } from 'react';
import { Container, Paper, Typography, Box, TextField, Button, Alert } from '@mui/material';
import { initiatePayment, getPaymentHistory } from '../services/api';

const Payments = () => {
  const [amount, setAmount] = useState('');
  const [method, setMethod] = useState('card');
  const [history, setHistory] = useState([]);
  const [message, setMessage] = useState('');

  useEffect(() => {
    const fetchHistory = async () => {
      try {
        const data = await getPaymentHistory();
        setHistory(data);
      } catch (err) {
        console.error(err);
      }
    };
    fetchHistory();
  }, []);

  const handlePay = async (e) => {
    e.preventDefault();
    try {
      const payload = { amount: Number(amount), method };
      const res = await initiatePayment(payload);
      setMessage('Payment initiated: ' + (res.id || 'success'));
      setAmount('');
      // refresh history
      const data = await getPaymentHistory();
      setHistory(data);
    } catch (err) {
      console.error(err);
      setMessage('Payment failed');
    }
  };

  return (
    <Container maxWidth="md">
      <Paper sx={{ p: 3, mt: 4 }}>
        <Typography variant="h6">Payments</Typography>
        {message && <Alert sx={{ mt: 2 }}>{message}</Alert>}
        <Box component="form" onSubmit={handlePay} sx={{ mt: 2, display: 'flex', gap: 2 }}>
          <TextField label="Amount" value={amount} onChange={(e) => setAmount(e.target.value)} required />
          <Button type="submit" variant="contained">Pay</Button>
        </Box>

        <Typography variant="h6" sx={{ mt: 3 }}>History</Typography>
        <Box sx={{ mt: 2 }}>
          {history.map((p) => (
            <Paper key={p.id} sx={{ p: 2, mb: 1 }}>
              <Typography>{p.id} - {p.amount}</Typography>
              <Typography variant="caption">{p.status}</Typography>
            </Paper>
          ))}
        </Box>
      </Paper>
    </Container>
  );
};

export default Payments;
