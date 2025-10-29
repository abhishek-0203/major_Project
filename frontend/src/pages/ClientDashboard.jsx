import React, { useEffect, useState } from 'react';
import {
  Container,
  Grid,
  Card,
  CardContent,
  Typography,
  Box,
  Button,
  Chip,
  CircularProgress,
  Paper,
  Divider,
  IconButton,
  useTheme
} from '@mui/material';
import {
  Add as AddIcon,
  Message as MessageIcon,
  Person as PersonIcon,
  Timeline as TimelineIcon
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';

const ClientDashboard = () => {
  const theme = useTheme();
  const navigate = useNavigate();
  const [projects, setProjects] = useState([]);
  const [developers, setDevelopers] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const token = localStorage.getItem('token');
        const headers = {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json'
        };

        // Fetch projects
        const projectsResponse = await fetch('http://localhost:8081/projects', { headers });
        const projectsData = await projectsResponse.json();

        // Fetch developers
        const developersResponse = await fetch('http://localhost:8081/developers', { headers });
        const developersData = await developersResponse.json();

        setProjects(projectsData);
        setDevelopers(developersData);
      } catch (error) {
        console.error('Error fetching dashboard data:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  const StatusChip = ({ status }) => {
    const getColor = () => {
      switch (status.toLowerCase()) {
        case 'open': return 'primary';
        case 'in progress': return 'warning';
        case 'completed': return 'success';
        default: return 'default';
      }
    };

    return (
      <Chip
        label={status}
        color={getColor()}
        size="small"
        sx={{ textTransform: 'capitalize' }}
      />
    );
  };

  if (loading) {
    return (
      <Box display="flex" justifyContent="center" alignItems="center" minHeight="80vh">
        <CircularProgress />
      </Box>
    );
  }

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      {/* Header Section */}
      <Box display="flex" justifyContent="space-between" alignItems="center" mb={4}>
        <Typography variant="h4" component="h1">
          Client Dashboard
        </Typography>
        <Button
          variant="contained"
          startIcon={<AddIcon />}
          onClick={() => navigate('/create-project')}
        >
          Post New Project
        </Button>
      </Box>

      <Grid container spacing={4}>
        {/* Statistics Cards */}
        <Grid item xs={12} md={4}>
          <Card sx={{ height: '100%', backgroundColor: theme.palette.primary.light }}>
            <CardContent>
              <Typography variant="h6" color="white" gutterBottom>
                Active Projects
              </Typography>
              <Typography variant="h3" color="white">
                {projects.filter(p => p.status === 'in progress').length}
              </Typography>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} md={4}>
          <Card sx={{ height: '100%', backgroundColor: theme.palette.secondary.light }}>
            <CardContent>
              <Typography variant="h6" color="white" gutterBottom>
                Completed Projects
              </Typography>
              <Typography variant="h3" color="white">
                {projects.filter(p => p.status === 'completed').length}
              </Typography>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} md={4}>
          <Card sx={{ height: '100%', backgroundColor: theme.palette.success.light }}>
            <CardContent>
              <Typography variant="h6" color="white" gutterBottom>
                Available Developers
              </Typography>
              <Typography variant="h3" color="white">
                {developers.filter(d => d.available === 'yes').length}
              </Typography>
            </CardContent>
          </Card>
        </Grid>

        {/* Projects List */}
        <Grid item xs={12}>
          <Paper sx={{ p: 3 }}>
            <Typography variant="h6" gutterBottom>
              Your Projects
            </Typography>
            <Grid container spacing={3}>
              {projects.map((project, index) => (
                <Grid item xs={12} key={index}>
                  <Card sx={{ 
                    p: 2,
                    '&:hover': {
                      boxShadow: 6,
                      transform: 'translateY(-2px)',
                      transition: 'all 0.3s'
                    }
                  }}>
                    <Grid container spacing={2} alignItems="center">
                      <Grid item xs={12} sm={6}>
                        <Typography variant="h6" gutterBottom>
                          {project.title}
                        </Typography>
                        <Typography variant="body2" color="text.secondary" gutterBottom>
                          {project.description}
                        </Typography>
                        <Box sx={{ mt: 1 }}>
                          <StatusChip status={project.status} />
                        </Box>
                      </Grid>
                      <Grid item xs={12} sm={3}>
                        <Box>
                          <Typography variant="subtitle2" color="text.secondary">
                            Budget
                          </Typography>
                          <Typography variant="h6">
                            ${project.budget}
                          </Typography>
                        </Box>
                      </Grid>
                      <Grid item xs={12} sm={3}>
                        <Box display="flex" justifyContent="flex-end" gap={1}>
                          <Button
                            variant="outlined"
                            size="small"
                            startIcon={<MessageIcon />}
                            onClick={() => navigate('/chat')}
                          >
                            Chat
                          </Button>
                          <Button
                            variant="contained"
                            size="small"
                            onClick={() => navigate(`/project/${project.id}`)}
                          >
                            View
                          </Button>
                        </Box>
                      </Grid>
                    </Grid>
                  </Card>
                </Grid>
              ))}
            </Grid>
          </Paper>
        </Grid>
      </Grid>
    </Container>
  );
};

export default ClientDashboard;