import React ,{useEffect, useState}from 'react';
// import Time from './Time';
// import './Postview.css';
import { Link } from 'react-router-dom';



const Postview=()=> {

    const [data ,setData] = useState([]);
    const getUserData = async () =>{
        const res = await fetch(process.env.REACT_APP_API+"/api/items")
        const posts = await res.json()
        console.log(posts);
        setData(posts)

      }


    useEffect(() => getUserData(), [])
    
    

  return (
    <>
    
    <div className="site-container">
    <div className="heading">
      <img
      style={{width:'30%',height:"100px",paddingRight:"300px"}} 
      src = "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTf2A64rqIzgBb2uqxicEOXOtjEUB9479dyeA&usqp=CAU"/>
      <Link to="/form" >
      {/* <img src = "https://thumbs.dreamstime.com/b/speed-fast-camera-icon-logo-design-element-can-be-used-as-as-complement-to-95291428.jpg"  className='txtHeader'/> */}
      <img src="https://img.icons8.com/nolan/96/shopping-cart-promotion.png"/>
      </Link>
    </div>
    <table>
    <thead>
            <tr>
              <th className="table-head ppdid" id="ppdid">
                Item
              </th>
              <th className="table-head image-head" id="image">
                Size
              </th>
              <th className="table-head" id="property">
                Colour
              </th>
              </tr>
            </thead>
            
    {data.map((dta,i) =>(
      <tr>
          <td className="item" key={i}>{dta.item}</td>
          <td className="size">{dta.size}</td>
          <td className="colour">{dta.colour}</td>
      </tr>
          ))}
    
      </table>
    </div>
    </>
  );
}

export default Postview;
