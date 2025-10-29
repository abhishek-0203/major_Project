import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Login from './pages/Login';
import Signup from './pages/Signup';
import Dashboard from './pages/Dashboard';
import ProjectForm from './pages/ProjectForm';
import Chat from './pages/Chat';
import Payments from './pages/Payments';
import ClientProfile from './pages/ClientProfile';
import DeveloperProfile from './pages/DeveloperProfile';
import ClientDashboard from './pages/ClientDashboard';
import DeveloperDashboard from './pages/DeveloperDashboard';
import Navbar from './components/Navbar';
import Home from './pages/Home';
import { ThemeProvider, Container } from '@mui/material';
import theme from './theme';

const PrivateRoute = ({ children }) => {
  const token = localStorage.getItem('token');
  const userRole = localStorage.getItem('userRole');
  
  if (!token) {
    return <Navigate to="/login" />;
  }
  
  // Role-based redirection for dashboard
  if (children.type === Dashboard) {
    return userRole === 'client' ? <ClientDashboard /> : <DeveloperDashboard />;
  }
  
  // Restrict project creation to clients
  if (children.type === ProjectForm && userRole !== 'client') {
    return <Navigate to="/dashboard" />;
  }
  
  return children;
};

function App() {
  return (
    <ThemeProvider theme={theme}>
      <Router>
        <Navbar />
        <Container sx={{ mt: 3 }}>
          <Routes>
            <Route path="/login" element={<Login />} />
            <Route path="/login/:role" element={<Login />} />
            <Route path="/signup" element={<Signup />} />
            <Route path="/signup/:role" element={<Signup />} />
            <Route path="/" element={<Home />} />
            <Route
              path="/dashboard"
              element={
                <PrivateRoute>
                  <Dashboard />
                </PrivateRoute>
              }
            />
            <Route
              path="/create-project"
              element={
                <PrivateRoute>
                  <ProjectForm />
                </PrivateRoute>
              }
            />
            <Route
              path="/chat"
              element={
                <PrivateRoute>
                  <Chat />
                </PrivateRoute>
              }
            />
            <Route
              path="/payments"
              element={
                <PrivateRoute>
                  <Payments />
                </PrivateRoute>
              }
            />
            <Route
              path="/profile/client"
              element={
                <PrivateRoute>
                  <ClientProfile />
                </PrivateRoute>
              }
            />
            <Route
              path="/profile/developer"
              element={
                <PrivateRoute>
                  <DeveloperProfile />
                </PrivateRoute>
              }
            />
            <Route path="/" element={<Navigate to="/login" />} />
          </Routes>
        </Container>
      </Router>
    </ThemeProvider>
  );
}

export default App;
