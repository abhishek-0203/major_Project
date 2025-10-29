import React from 'react';
import { Box, Button, Typography } from '@mui/material';
import RoleSelection from '../components/RoleSelection';

const Home = () => {
  return (
    <div className="container">
      <section className="hero">
        <div>
          <Typography component="h1">DevConnect — Hire, Collaborate, Deliver</Typography>
          <Typography className="muted mt-2">A marketplace built for quality and trust — connect with top developers and securely manage payments.</Typography>
          <Box sx={{ mt: 3 }}>
            <Button variant="contained" sx={{ mr: 2 }} href="/signup/client">I'm a client</Button>
            <Button variant="outlined" href="/signup/developer">I'm a developer</Button>
          </Box>
        </div>
        <img src="/devconnect-logo.svg" alt="DevConnect" style={{ maxWidth: 240, opacity: 0.95 }} />
      </section>

      <section style={{ marginTop: 32 }}>
        <Typography variant="h5" sx={{ mb: 2 }}>Get started</Typography>
        <RoleSelection mode="signup" />
      </section>
    </div>
  );
};

export default Home;
