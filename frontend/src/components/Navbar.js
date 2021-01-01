import React from 'react'
import { Link } from 'react-router-dom';
import './Navbar.css'

function Navbar() {
    return (
        <div>
            <nav className="navbar">
                <div className="navbar-container">
                    <Link to="/" className="navbar-logo">
                        PicStore <i className = 'fab fa-typo3' />   
                    </Link>
                </div>
            </nav>
        </div>
    )
}

export default Navbar
