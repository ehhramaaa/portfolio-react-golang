
const Navbar = () => {
    return (
        <div className="container pl-0 pr-0">
            <div className="header-con">
                <nav className="navbar navbar-expand-lg navbar-light p-0">
                    <a className="navbar-brand p-0" href="">
                        <img src="/logo.png" alt="logo-img" className="img-fluid w-50"/>
                    </a>
                    <button className="navbar-toggler p-0 collapsed" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                        <span className="navbar-toggler-icon" />
                        <span className="navbar-toggler-icon" />
                        <span className="navbar-toggler-icon" />
                    </button>
                    <div className="collapse navbar-collapse justify-content-end" id="navbarSupportedContent">
                        <div className="d-inline-block contact">
                            <a href="#Contact">Contact</a>
                        </div>
                    </div>
                </nav>
            </div>
        </div>

    )
}

export default Navbar