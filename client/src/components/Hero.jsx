const Hero = () => {
    return (
        <>
            <div className="footer-social-icon banner-social-icon mb-0">
                <ul className="mb-0 list-unstyled">
                    <li className>
                        <a href="https://github.com/ehhramaaa" target="_blank"><i className="fab fa-github d-flex align-items-center justify-content-center" /></a>
                    </li>
                    <li className="mt-3 mb-3">
                        <a href="https://www.instagram.com/ehhramaa_/" target="_blank"><i className="fab fa-instagram d-flex align-items-center justify-content-center ml-0 mr-0 " /></a>
                    </li>
                    <li className>
                        <a href="https://www.linkedin.com/in/muhamad-ramadhan-bb6289237/" target="_blank"><i className="fab fa-linkedin-in d-flex align-items-center justify-content-center" /></a>
                    </li>
                </ul>
            </div>
            <div className="banner-con text-lg-left text-center">
                <div className="row">
                    <div className="col-lg-7 col-sm-12 d-flex justify-content-center flex-column">
                        <div className="banner-left-con wow slideInLeft">
                            <div className="banner-heading">
                                <h2>Hello, I Am</h2>
                                <ul className="dynamic-txts">
                                    <li><h1>Muhamad Ramadhan</h1></li>
                                    <li><h1>Muhamad Ramadhan</h1></li>
                                    <li><h1>Muhamad Ramadhan</h1></li>
                                    <li><h1>Muhamad Ramadhan</h1></li>
                                </ul>
                                <p>5th semester Information Systems student who has a strong passion for backend development.<br />
                                </p>
                            </div>
                            <div className="banner-btn generic-btn d-inline-block">
                                <a href="#Contact">More About Me</a>
                            </div>
                            <a href="#Portfolio" className="See-btn">See My Work</a>
                        </div>
                    </div>
                    <div className="col-lg-5 col-sm-12">
                        <div className="banner-right-con position-relative wow slideInRight" id="banner-right-con">
                            <figure className="mb-0">
                                <img src="/me.png" alt="banner-right-img" id="banner-right-img" className="w-100" />
                            </figure>
                            <div className="best-award-con d-inline-block wow bounceInUp" data-wow-duration="1s" data-wow-delay="1s">
                                <div className="best-award-inner-con">
                                    <figure className="mb-0">
                                        <img src="/experience_1.png" alt="cup-img" className="img-fluid" />
                                    </figure>
                                    <div className="best-award-title">
                                        <p className="mb-0">6+ Project<br />
                                            Completed.
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="best-award-con d-inline-block happy-con wow bounceInUp " data-wow-duration="1s" data-wow-delay="1s">
                                <div className="best-award-inner-con text-center">
                                    <figure>
                                        <img src="/image/admin-icon.png" alt="admin-icon" className="img-fluid" />
                                    </figure>
                                    <div className="best-award-title d-inline-block ml-0">
                                        <p className="mb-0 d-inline-block count">6</p>
                                        <p className="mb-0 d-inline-block">+</p>
                                        <span className="d-block">Month<br />
                                            Experience</span>
                                    </div>
                                </div>
                            </div>
                            <div className="cursor" />
                            <div className="cursor2" />
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}

export default Hero