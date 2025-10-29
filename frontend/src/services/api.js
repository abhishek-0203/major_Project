import axios from 'axios';

const API_URL = 'http://localhost:8081';

const api = axios.create({
  baseURL: API_URL,
});

// Add a request interceptor to include the JWT token
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Auth
export const login = async (email, password, role) => {
  const payload = { email, password };
  if (role) payload.role = role;
  const response = await api.post('/login', payload);
  return response.data;
};

export const signup = async (userData) => {
  const response = await api.post('/signup', userData);
  return response.data;
};

// Projects
export const getProjects = async () => {
  const response = await api.get('/projects');
  return response.data;
};

export const createProject = async (projectData) => {
  const response = await api.post('/createProject', projectData);
  return response.data;
};

export const updateProject = async (projectId, projectData) => {
  const response = await api.put(`/updateProject/${projectId}`, projectData);
  return response.data;
};

export const deleteProject = async (projectId) => {
  const response = await api.delete(`/deleteProject/${projectId}`);
  return response.data;
};

// Profiles (clients & developers)
export const getClients = async () => {
  const response = await api.get('/clients');
  return response.data;
};

export const createClientProfile = async (profileData) => {
  const response = await api.post('/createClientProfile', profileData);
  return response.data;
};

export const updateClientProfile = async (index, profileData) => {
  const response = await api.put(`/updateClientProfile/${index}`, profileData);
  return response.data;
};

export const deleteClientProfile = async (index) => {
  const response = await api.delete(`/deleteClientProfile/${index}`);
  return response.data;
};

export const getDevelopers = async () => {
  const response = await api.get('/developers');
  return response.data;
};

export const createDeveloperProfile = async (profileData) => {
  const response = await api.post('/createDeveloperProfile', profileData);
  return response.data;
};

export const updateDeveloperProfile = async (index, profileData) => {
  const response = await api.put(`/updateDeveloperProfile/${index}`, profileData);
  return response.data;
};

export const deleteDeveloperProfile = async (index) => {
  const response = await api.delete(`/deleteDeveloperProfile/${index}`);
  return response.data;
};

// Payments
export const initiatePayment = async (paymentData) => {
  const response = await api.post('/api/payment/initiate', paymentData);
  return response.data;
};

export const getPaymentHistory = async () => {
  const response = await api.get('/api/payment/all');
  return response.data;
};

export const getPaymentDetails = async (paymentId) => {
  const response = await api.get(`/api/payment/${paymentId}`);
  return response.data;
};

export default api;