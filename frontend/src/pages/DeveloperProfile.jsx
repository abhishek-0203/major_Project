import React, { useEffect, useState } from 'react';
import { Container, Paper, Typography, TextField, Button, Alert } from '@mui/material';
import { getDevelopers, updateDeveloperProfile } from '../services/api';

const DeveloperProfile = () => {
  const [profile, setProfile] = useState({});
  const [msg, setMsg] = useState('');

  useEffect(() => {
    const fetch = async () => {
      try {
        const data = await getDevelopers();
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
      await updateDeveloperProfile(profile.index || 0, profile);
      setMsg('Profile updated');
    } catch (err) {
      console.error(err);
      setMsg('Update failed');
    }
  };

  return (
    <Container maxWidth="sm">
      <Paper sx={{ p: 3, mt: 4 }}>
        <Typography variant="h6">Developer Profile</Typography>
        {msg && <Alert sx={{ mt: 2 }}>{msg}</Alert>}
        <TextField label="Name" name="name" value={profile.name || ''} onChange={handleChange} fullWidth sx={{ mt: 2 }} />
        <TextField label="Email" name="email" value={profile.email || ''} onChange={handleChange} fullWidth sx={{ mt: 2 }} />
        <TextField label="Skills" name="skills" value={profile.skills || ''} onChange={handleChange} fullWidth sx={{ mt: 2 }} />
        <Button variant="contained" sx={{ mt: 2 }} onClick={handleSave}>Save</Button>
      </Paper>
    </Container>
  );
};

export default DeveloperProfile;
