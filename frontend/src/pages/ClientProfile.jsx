import React, { useEffect, useState } from 'react';
import { Container, Paper, Typography, TextField, Button, Alert } from '@mui/material';
import { getClients, updateClientProfile } from '../services/api';

const ClientProfile = () => {
  const [profile, setProfile] = useState({});
  const [msg, setMsg] = useState('');

  useEffect(() => {
    const fetch = async () => {
      try {
        const data = await getClients();
        // backend returns array; pick first or find by id in a real app
        setProfile(data[0] || {});
      } catch (err) {
        console.error(err);
      }
    };
    fetch();
  }, []);

  const handleChange = (e) => {
    setProfile({ ...profile, [e.target.name]: e.target.value });
  };

  const handleSave = async () => {
    try {
      // using index 0 for example; adapt to real id
      await updateClientProfile(profile.index || 0, profile);
      setMsg('Profile updated');
    } catch (err) {
      console.error(err);
      setMsg('Update failed');
    }
  };

  return (
    <Container maxWidth="sm">
      <Paper sx={{ p: 3, mt: 4 }}>
        <Typography variant="h6">Client Profile</Typography>
        {msg && <Alert sx={{ mt: 2 }}>{msg}</Alert>}
        <TextField label="Name" name="name" value={profile.name || ''} onChange={handleChange} fullWidth sx={{ mt: 2 }} />
        <TextField label="Email" name="email" value={profile.email || ''} onChange={handleChange} fullWidth sx={{ mt: 2 }} />
        <TextField label="Company" name="company" value={profile.company || ''} onChange={handleChange} fullWidth sx={{ mt: 2 }} />
        <Button variant="contained" sx={{ mt: 2 }} onClick={handleSave}>Save</Button>
      </Paper>
    </Container>
  );
};

export default ClientProfile;
