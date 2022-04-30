import { BrowserRouter, Route, Routes } from "react-router-dom";
import GAbout from "./components/GAbout";
import GContainer from "./components/GContainer";
import GHome from "./components/GHome";
import GLogin from "./components/GLogin";
import GLogout from "./components/GLogout";
import GPageNotFound from "./components/GPageNotFound";
import GPasteSearch from "./components/GPasteSearch";
import GPasteSearchContent from "./components/GPasteSearchContent";
import GRegister from "./components/GRegister";
import GUser from "./components/GUserManager";
import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<GContainer />} >
          <Route index element={<GHome />} />
          <Route path="register" element={<GRegister />} />
          <Route path="login" element={<GLogin />} />
          <Route path="logout" element={<GLogout />} />
          <Route path="user" element={<GUser />} />
          <Route path="about" element={<GAbout />} />
          <Route path="s" element={<GPasteSearch />}>
            <Route path=":key" element={<GPasteSearchContent />} />
          </Route>
          <Route path="*" element={<GPageNotFound />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App;