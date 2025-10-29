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
    <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        gap: 4,
        p: 4,
      }}
    >
      <Typography variant="h4" gutterBottom>
        Select Your Role
      </Typography>
      <Box sx={{ display: 'flex', gap: 4 }}>
        <Paper
          elevation={3}
          sx={{
            p: 3,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            cursor: 'pointer',
            '&:hover': { bgcolor: 'action.hover' },
          }}
          onClick={() => handleRoleSelect('client')}
        >
          <Person sx={{ fontSize: 60, mb: 2 }} />
          <Typography variant="h6">Client</Typography>
          <Typography variant="body2" color="text.secondary" align="center">
            Looking to hire developers
          </Typography>
        </Paper>

        <Paper
          elevation={3}
          sx={{
            p: 3,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            cursor: 'pointer',
            '&:hover': { bgcolor: 'action.hover' },
          }}
          onClick={() => handleRoleSelect('developer')}
        >
          <Code sx={{ fontSize: 60, mb: 2 }} />
          <Typography variant="h6">Developer</Typography>
          <Typography variant="body2" color="text.secondary" align="center">
            Looking for projects
          </Typography>
        </Paper>
      </Box>
    </Box>
  );
};

export default RoleSelection;