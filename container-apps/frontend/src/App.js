import { Outlet } from "react-router-dom"
import Header from "./components/Header";

function App() {
  return (
    <div>
      <Header />
      <hr />
      <Outlet />
    </div>
  )
};

export default App;
