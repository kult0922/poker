import { BrowserRouter, Route } from 'react-router-dom';
import { Home } from './pages/home';
import { HandDemo } from './pages/hand-demo';

function App() {
  return (
    <BrowserRouter>
      <Route exact path="/">
        <Home />
      </Route>
      <Route path="/hand-demo">
        <HandDemo />
      </Route>
    </BrowserRouter>
  );

  return <div>failed to load</div>;
}

export default App;
