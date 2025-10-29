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
  Avatar,
  Rating,
  useTheme
} from '@mui/material';
import {
  Work as WorkIcon,
  Message as MessageIcon,
  AttachMoney as MoneyIcon,
  Star as StarIcon
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';

const DeveloperDashboard = () => {
  const theme = useTheme();
  const navigate = useNavigate();
  const [projects, setProjects] = useState([]);
  const [activeProjects, setActiveProjects] = useState([]);
  const [earnings, setEarnings] = useState(0);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const token = localStorage.getItem('token');
        const headers = {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json'
        };

        // Fetch available projects
        const projectsResponse = await fetch('http://localhost:8081/projects', { headers });
        const projectsData = await projectsResponse.json();
        setProjects(projectsData);

        // Fetch active projects
        // This would be developer-specific projects in a real app
        setActiveProjects(projectsData.filter(p => p.status === 'in progress'));

        // Fetch earnings (mock data for now)
        setEarnings(5000);
      } catch (error) {
        console.error('Error fetching dashboard data:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

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
      <Grid container spacing={4}>
        {/* Statistics Cards */}
        <Grid item xs={12} md={4}>
          <Card sx={{ 
            height: '100%',
            background: `linear-gradient(45deg, ${theme.palette.primary.main}, ${theme.palette.primary.light})`
          }}>
            <CardContent>
              <Box display="flex" alignItems="center" mb={2}>
                <WorkIcon sx={{ color: 'white', mr: 1 }} />
                <Typography variant="h6" color="white">
                  Active Projects
                </Typography>
              </Box>
              <Typography variant="h3" color="white">
                {activeProjects.length}
              </Typography>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} md={4}>
          <Card sx={{ 
            height: '100%',
            background: `linear-gradient(45deg, ${theme.palette.secondary.main}, ${theme.palette.secondary.light})`
          }}>
            <CardContent>
              <Box display="flex" alignItems="center" mb={2}>
                <MoneyIcon sx={{ color: 'white', mr: 1 }} />
                <Typography variant="h6" color="white">
                  Total Earnings
                </Typography>
              </Box>
              <Typography variant="h3" color="white">
                ${earnings}
              </Typography>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} md={4}>
          <Card sx={{ 
            height: '100%',
            background: `linear-gradient(45deg, ${theme.palette.success.main}, ${theme.palette.success.light})`
          }}>
            <CardContent>
              <Box display="flex" alignItems="center" mb={2}>
                <StarIcon sx={{ color: 'white', mr: 1 }} />
                <Typography variant="h6" color="white">
                  Rating
                </Typography>
              </Box>
              <Box display="flex" alignItems="center">
                <Typography variant="h3" color="white" mr={1}>
                  4.8
                </Typography>
                <Rating value={4.8} readOnly sx={{ color: 'white' }} />
              </Box>
            </CardContent>
          </Card>
        </Grid>

        {/* Available Projects */}
        <Grid item xs={12}>
          <Paper sx={{ p: 3 }}>
            <Typography variant="h6" gutterBottom>
              Available Projects
            </Typography>
            <Grid container spacing={3}>
              {projects.filter(p => p.status === 'open').map((project, index) => (
                <Grid item xs={12} key={index}>
                  <Card sx={{ 
                    p: 2,
                    '&:hover': {
                      boxShadow: 6,
                      transform: 'translateY(-2px)',
                      transition: 'all 0.3s'
                    }
                  }}>
                    <Grid container spacing={2}>
                      <Grid item xs={12} sm={6}>
                        <Typography variant="h6" gutterBottom>
                          {project.title}
                        </Typography>
                        <Typography variant="body2" color="text.secondary" paragraph>
                          {project.description}
                        </Typography>
                        <Box display="flex" flexWrap="wrap" gap={1}>
                          {project.requirements?.map((skill, i) => (
                            <Chip
                              key={i}
                              label={skill}
                              size="small"
                              color="primary"
                              variant="outlined"
                            />
                          ))}
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
                          <Typography variant="subtitle2" color="text.secondary" mt={1}>
                            Timeline
                          </Typography>
                          <Typography variant="body1">
                            {project.timeline}
                          </Typography>
                        </Box>
                      </Grid>
                      <Grid item xs={12} sm={3}>
                        <Box display="flex" flexDirection="column" gap={1}>
                          <Button
                            variant="contained"
                            fullWidth
                            onClick={() => navigate(`/project/${project.id}`)}
                          >
                            Apply Now
                          </Button>
                          <Button
                            variant="outlined"
                            fullWidth
                            startIcon={<MessageIcon />}
                            onClick={() => navigate('/chat')}
                          >
                            Contact Client
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

export default DeveloperDashboard;