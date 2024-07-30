const Service = () => {
    return (
        <section className="w-100 float-left service-con padding-top padding-bottom position-relative" id="service-con">
            <div className="container">
                <div className="service-inner-con position-relative">
                    <div className="generic-title text-center">
                        <h6>My Expertise</h6>
                        <h2 className="mb-0">Provide Wide Range of<br />
                            Digital Services
                        </h2>
                    </div>
                    <div className="service-box wow fadeInUp">
                        <div className="row mb-0">
                            <div className="col-lg-6 col-md-6">
                                <div className="service-box-item h-100">
                                    <figure className="mb-0">
                                        <img src="/image/service-icon3.png" alt="service-icon" className="img-fluid" />
                                    </figure>
                                    <div className="service-box-item-content d-flex flex-column justify-content-between">
                                        <h4>Web Development</h4>
                                        <p>
                                            Im ready to help you realize your digital vision. My focus is on creating engaging user experiences and optimal performance.
                                        </p>
                                        <a href="#" data-toggle="modal" data-target="#web-development">Read More</a>
                                    </div>
                                </div>
                            </div>
                            <div className="col-lg-6 col-md-6">
                                <div className="service-box-item mb-0 h-100">
                                    <figure className="mb-0">
                                        <img src="/image/service-icon4.png" alt="service-icon" className="img-fluid" />
                                    </figure>
                                    <div className="service-box-item-content d-flex flex-column justify-content-between">
                                        <h4>App Development</h4>
                                        <p>
                                            I can create applications that are functional, user-friendly, and suit your business needs.
                                        </p>
                                        <a href="#" data-toggle="modal" data-target="#app-development">Read More</a>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    )
}

export default Service