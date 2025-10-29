import React from 'react';
import { AppBar, Toolbar, Typography, Button, Box, IconButton } from '@mui/material';
import { useNavigate } from 'react-router-dom';
// logo is served from public/devconnect-logo.svg

const Navbar = () => {
  const navigate = useNavigate();
  const role = localStorage.getItem('userRole');

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('userRole');
    navigate('/login');
  };

  return (
    <AppBar position="static" sx={{ background: 'linear-gradient(90deg,#07102a, #0b1b33)' }}>
      <Toolbar sx={{ display: 'flex', gap: 2 }}>
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 1, cursor: 'pointer' }} onClick={() => navigate('/')}>
          <img src="/devconnect-logo.svg" alt="DevConnect" style={{ width: 36, height: 36, borderRadius: 8 }} />
          <Typography variant="h6" sx={{ fontWeight: 700 }}>DevConnect</Typography>
        </Box>

        <Box sx={{ display: 'flex', gap: 1, alignItems: 'center' }}>
          {role === 'client' && (
            <>
              <Button color="inherit" variant="outlined" onClick={() => navigate('/create-project')}>Create Project</Button>
              <Button color="inherit" onClick={() => navigate('/profile/client')}>My Profile</Button>
            </>
          )}
          {role === 'developer' && (
            <Button color="inherit" onClick={() => navigate('/profile/developer')}>My Profile</Button>
          )}
          <Button color="inherit" onClick={() => navigate('/chat')}>Chat</Button>
          <Button color="inherit" onClick={() => navigate('/payments')}>Payments</Button>
          <Button color="secondary" sx={{ ml: 1 }} onClick={handleLogout}>Logout</Button>
        </Box>
      </Toolbar>
    </AppBar>
  );
};

export default Navbar;
