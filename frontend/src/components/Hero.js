import React from 'react';
import blob1 from '../assets/blobs/blob1.svg';
//import blob2 from '../assets/blobs/blob2.svg';
import './Hero.css';
function Hero() {
    return (
        <div className= 'hero-container'>
            <div className = "hero-img-blob" >
                <img
                    src={blob1}
                    alt='hero'
                    className= "hero-img-blob1"
                    
                />
                {/* <img
                    src={blob2}
                    alt='hero'
                    height='230vw'
                    className= "hero-img-blob2"
                /> */}
            </div>
            
        </div>
    )
}

export default Hero
