import React from 'react';
import { Box, Button, Typography, Paper } from '@mui/material';
import { Person, Code } from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';

const RoleSelection = ({ mode }) => {
  const navigate = useNavigate();

  const handleRoleSelect = (role) => {
    navigate(`/${mode}/${role}`);
  };

  return (
    <Box sx={{ display: 'flex', justifyContent: 'center', p: 4 }}>
      <Box sx={{ display: 'flex', gap: 24, alignItems: 'stretch', flexWrap: 'wrap', justifyContent: 'center' }}>
        <Paper
          className="card"
          elevation={6}
          sx={{ width: 320, p: 4, display: 'flex', flexDirection: 'column', alignItems: 'center', cursor: 'pointer' }}
          onClick={() => handleRoleSelect('client')}
        >
          <Box sx={{ width: 72, height: 72, borderRadius: 16, display: 'grid', placeItems: 'center', background: 'linear-gradient(135deg,#6a9cff,#7b61ff)', mb: 2 }}>
            <Person sx={{ color: 'white', fontSize: 36 }} />
          </Box>
          <Typography variant="h6">Client</Typography>
          <Typography className="muted" sx={{ textAlign: 'center', mt: 1 }}>Hire pre-vetted developers for your projects.</Typography>
          <Button variant="contained" sx={{ mt: 3, width: '100%' }} onClick={() => handleRoleSelect('client')}>Get Started</Button>
        </Paper>

        <Paper
          className="card"
          elevation={6}
          sx={{ width: 320, p: 4, display: 'flex', flexDirection: 'column', alignItems: 'center', cursor: 'pointer' }}
          onClick={() => handleRoleSelect('developer')}
        >
          <Box sx={{ width: 72, height: 72, borderRadius: 16, display: 'grid', placeItems: 'center', background: 'linear-gradient(135deg,#34d399,#3b82f6)', mb: 2 }}>
            <Code sx={{ color: 'white', fontSize: 36 }} />
          </Box>
          <Typography variant="h6">Developer</Typography>
          <Typography className="muted" sx={{ textAlign: 'center', mt: 1 }}>Find meaningful projects and grow your portfolio.</Typography>
          <Button variant="contained" sx={{ mt: 3, width: '100%' }} onClick={() => handleRoleSelect('developer')}>Join Now</Button>
        </Paper>
      </Box>
    </Box>
  );
};

export default RoleSelection;