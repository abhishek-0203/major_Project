import React, { useEffect, useState } from 'react';
import { Box, Container, Grid, Card, CardContent, Typography, CardActions, Button, Avatar } from '@mui/material';
import { getProjects } from '../services/api';

const ProjectCard = ({ project }) => (
  <Card sx={{ minHeight: 180, background: 'transparent', border: '1px solid rgba(255,255,255,0.04)' }}>
    <CardContent>
      <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
        <Avatar sx={{ bgcolor: 'primary.main' }}>{(project.title || 'P')[0]}</Avatar>
        <Box>
          <Typography variant="h6">{project.title}</Typography>
          <Typography className="muted" variant="body2">{project.owner || 'Unknown'} · ${project.budget || '—'}</Typography>
        </Box>
      </Box>
      <Typography sx={{ mt: 2 }} className="muted">{project.description}</Typography>
    </CardContent>
    <CardActions>
      <Button size="small">View</Button>
      <Button size="small" variant="contained">Apply</Button>
    </CardActions>
  </Card>
);

const Dashboard = () => {
  const [projects, setProjects] = useState([]);

  useEffect(() => {
    const fetchProjects = async () => {
      try {
        const data = await getProjects();
        setProjects(data || []);
      } catch (error) {
        console.error('Error fetching projects:', error);
      }
    };

    fetchProjects();
  }, []);

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Box className="hero">
        <Box>
          <Typography variant="h4">Find great projects or hire top developers</Typography>
          <Typography className="muted mt-2">Browse vetted projects, post your own, and collaborate with confidence.</Typography>
          <Box sx={{ mt: 3 }}>
            <Button variant="contained" sx={{ mr: 2 }} href="/create-project">Create Project</Button>
            <Button variant="outlined" href="/chat">Join Chat</Button>
          </Box>
        </Box>
        <img src="/hero-illustration.png" alt="hero" style={{ maxWidth: 420, opacity: 0.9 }} />
      </Box>

      <Grid container spacing={3} sx={{ mt: 2 }}>
        {projects.length === 0 && (
          <Grid item xs={12}>
            <Box className="card" sx={{ p: 3 }}>
              <Typography className="muted">No projects yet. Clients can create a project using the Create Project button.</Typography>
            </Box>
          </Grid>
        )}

        {projects.map((project, idx) => (
          <Grid item xs={12} sm={6} md={4} key={project.id || idx}>
            <ProjectCard project={project} />
          </Grid>
        ))}
      </Grid>
    </Container>
  );
};

export default Dashboard;