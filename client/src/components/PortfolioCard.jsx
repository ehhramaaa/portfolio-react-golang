const PortfolioCard = () => {
    return (
        <section className="w-100 float-left portfolio-body-con position-relative">
            <div className="container">
                <div className="portfolio-img-con position-relative w-100 float-left wow fadeInUp">
                    <div className="filterDiv cars position-relative">
                        <a href="#" data-toggle="modal" data-target="#modalWPWAF">
                            <div className="portfolio-img position-relative">
                                <figure>
                                    <img src="assets/image/portfolio-img1.png" alt="portfolio-img1" className="img-fluid" />
                                </figure>
                            </div>
                        </a>
                        <div className="portfolio-img-content text-left">
                            <div className="portfolio-img-title d-inline-block">
                                <h4>Application UI Design</h4>
                                <p>Dolar repellendus temporibus...</p>
                            </div>
                            <a href="#" className="float-lg-right" data-toggle="modal" data-target="#modalWPWAF-icon">
                                <i className="fas fa-arrow-right d-flex align-items-center justify-content-center" />
                            </a>
                        </div>
                    </div>
                    <div className="filterDiv colors fruits position-relative">
                        <a href="#" data-toggle="modal" data-target="#modalporfolio2">
                            <div className="portfolio-img position-relative">
                                <figure>
                                    <img src="assets/image/portfolio-img2.png" alt="portfolio-img1" className="img-fluid" />
                                </figure>
                            </div>
                        </a>
                        <div className="portfolio-img-content text-left">
                            <div className="portfolio-img-title d-inline-block">
                                <h4>Furni furniture UI Design</h4>
                                <p>Dolar repellendus temporibus...</p>
                            </div>
                            <a href="#" className="float-lg-right" data-toggle="modal" data-target="#modalporfolio2-icon">
                                <i className="fas fa-arrow-right d-flex align-items-center justify-content-center" />
                            </a>
                        </div>
                    </div>
                    <div className="filterDiv cars position-relative">
                        <a href="#" data-toggle="modal" data-target="#modalporfolio3">
                            <div className="portfolio-img position-relative">
                                <figure>
                                    <img src="assets/image/portfolio-img3.png" alt="portfolio-img1" className="img-fluid" />
                                </figure>
                            </div>
                        </a>
                        <div className="portfolio-img-content text-left">
                            <div className="portfolio-img-title d-inline-block">
                                <h4>Mobile UI Design</h4>
                                <p>Dolar repellendus temporibus...</p>
                            </div>
                            <a href="#" className="float-lg-right" data-toggle="modal" data-target="#modalporfolio3-icon"><i className="fas fa-arrow-right d-flex align-items-center justify-content-center" /></a>
                        </div>
                    </div>
                    <div className="filterDiv colors position-relative">
                        <a href="#" data-toggle="modal" data-target="#modalporfolio4">
                            <div className="portfolio-img position-relative">
                                <figure>
                                    <img src="assets/image/portfolio-img4.png" alt="portfolio-img1" className="img-fluid" />
                                </figure>
                            </div>
                        </a>
                        <div className="portfolio-img-content text-left">
                            <div className="portfolio-img-title d-inline-block">
                                <h4>Businesscard UI Design</h4>
                                <p>Dolar repellendus temporibus...</p>
                            </div>
                            <a href="#" className="float-lg-right" data-toggle="modal" data-target="#modalporfolio4-icon"><i className="fas fa-arrow-right d-flex align-items-center justify-content-center" /></a>
                        </div>
                    </div>
                    <div className="filterDiv cars animals position-relative">
                        <a href="#" data-toggle="modal" data-target="#modalporfolio5">
                            <div className="portfolio-img position-relative">
                                <figure>
                                    <img src="assets/image/portfolio-img5.png" alt="portfolio-img1" className="img-fluid" />
                                </figure>
                            </div>
                        </a>
                        <div className="portfolio-img-content text-left">
                            <div className="portfolio-img-title d-inline-block">
                                <h4>Real estate UI Design</h4>
                                <p>Dolar repellendus temporibus...</p>
                            </div>
                            <a href="#" className="float-lg-right" data-toggle="modal" data-target="#modalporfolio5-icon"><i className="fas fa-arrow-right d-flex align-items-center justify-content-center" /></a>
                        </div>
                    </div>
                </div>
            </div>
        </section>

    )
}

export default PortfolioCard