import React, { useState } from 'react';
import {
  Box,
  Button,
  TextField,
  Typography,
  Container,
  Alert,
  MenuItem,
  Paper,
} from '@mui/material';
import { createProject } from '../services/api';

const ProjectForm = () => {
  const [formData, setFormData] = useState({
    title: '',
    description: '',
    budget: '',
    duration: '',
    skills: '',
    projectType: 'fixed',
  });
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await createProject(formData);
      setSuccess('Project created successfully!');
      setFormData({
        title: '',
        description: '',
        budget: '',
        duration: '',
        skills: '',
        projectType: 'fixed',
      });
    } catch (err) {
      setError(err.response?.data?.message || 'An error occurred');
    }
  };

  return (
    <Container component="main" maxWidth="md">
      <Paper elevation={3} sx={{ p: 4, mt: 4 }}>
        <Typography component="h1" variant="h5" gutterBottom>
          Create New Project
        </Typography>
        {error && <Alert severity="error" sx={{ mb: 2 }}>{error}</Alert>}
        {success && <Alert severity="success" sx={{ mb: 2 }}>{success}</Alert>}
        <Box component="form" onSubmit={handleSubmit}>
          <TextField
            margin="normal"
            required
            fullWidth
            id="title"
            label="Project Title"
            name="title"
            value={formData.title}
            onChange={handleChange}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            multiline
            rows={4}
            id="description"
            label="Project Description"
            name="description"
            value={formData.description}
            onChange={handleChange}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            id="budget"
            label="Budget (USD)"
            name="budget"
            type="number"
            value={formData.budget}
            onChange={handleChange}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            id="duration"
            label="Duration (in weeks)"
            name="duration"
            type="number"
            value={formData.duration}
            onChange={handleChange}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            id="skills"
            label="Required Skills (comma-separated)"
            name="skills"
            value={formData.skills}
            onChange={handleChange}
            helperText="Example: React, Node.js, MongoDB"
          />
          <TextField
            margin="normal"
            required
            fullWidth
            select
            id="projectType"
            label="Project Type"
            name="projectType"
            value={formData.projectType}
            onChange={handleChange}
          >
            <MenuItem value="fixed">Fixed Price</MenuItem>
            <MenuItem value="hourly">Hourly Rate</MenuItem>
          </TextField>
          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >
            Create Project
          </Button>
        </Box>
      </Paper>
    </Container>
  );
};

export default ProjectForm;