import "./Form.css"
import {useState} from "react";
import { useNavigate } from 'react-router-dom';
import axios from 'axios';

const Form = () => {
    const [user, setUser] = useState({
        item: '',
        size: '',
        colour: ''
      });
      const navigate = useNavigate();
    
      const handleSubmit = async (e) => {
        
        e.preventDefault();
        const formData = new FormData();
    
        // const image = e.target.elements.userFile.files[0];
    
        // formData.append('image', image);
        formData.append('item', user.author);
        formData.append('size', user.location);
        formData.append('colour', user.description);
        
    
        try {
       
          const response = await axios.post(
            process.env.REACT_APP_API + '/api/item',
            formData
          );
          navigate('/');
          console.log(response);
        } catch (error) {
          console.log(error);
        }
      };
    return (  
        <>
        <div className='form-wrapper'>
        <div className='form'>
            <form id="#main-div"  onSubmit={handleSubmit}>
            <div className='item'>
                <input 
                type="text"
                placeholder="Item"
                value = {user.item}
                onChange={(e) => setUser({ ...user, item: e.target.value })}
                required/>
                <input type="text"
                 placeholder="Size" 
                 value = {user.size}
                 onChange={(e) => setUser({ ...user, size: e.target.value })}
                 required/>
               </div>
               <div className='desc'>
                <input type="text"
                 placeholder="Colour"
                 value = {user.colour}
                 onChange={(e) => setUser({ ...user, colour: e.target.value })}
                 />
                </div>
                <div className='post'>
                <button id="form-btn" type='submit'>Post</button>
                </div>
            </form>
          </div>
        </div>
        </>
    );
}
 
export default Form;