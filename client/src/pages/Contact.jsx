const Contact = () => {
    return (
        <section className="w-100 float-left form-main-con padding-top padding-bottom" id="Contact">
            <div className="container">
                <div className="form-main-inner-con position-relative">
                    <div className="generic-title text-center">
                        <h6>Get in Touch</h6>
                        <h2 className="mb-0">Any Questions? Feel Free<br />
                            to Contact
                        </h2>
                    </div>
                    <div className="row">
                        <div className="col-lg-4 order-lg-0 order-2">
                            <div className="contact-information position-relative wow slideInLeft">
                                <ul className="list-unstyled">
                                    <li>
                                        <figure className="mb-0 d-flex align-items-center justify-content-center">
                                            <img src="/image/location-icon.png" alt="location-icon" className="img-fluid" />
                                        </figure>
                                        <div className="contact-information-content">
                                            <h5>Address:</h5>
                                            <p className="mb-0">
                                                Jl. Cikutra Baru Raya,
                                                Kota Bandung, Jawa Barat
                                            </p>
                                        </div>
                                    </li>
                                    <li>
                                        <figure className="mb-0 d-flex align-items-center justify-content-center">
                                            <img src="/image/message-icon.png" alt="message-icon" className="img-fluid" />
                                        </figure>
                                        <div className="contact-information-content">
                                            <h5>Email:</h5>
                                            <p className="mb-0">ehhramaa212@gmail.com</p>
                                        </div>
                                    </li>
                                    <li className="mb-0">
                                        <figure className="mb-0 d-flex align-items-center justify-content-center">
                                            <img src="/image/phone-icon.png" alt="phone-icon" className="img-fluid" />
                                        </figure>
                                        <div className="contact-information-content">
                                            <h5>Phone:</h5>
                                            <p className="mb-0">+6281317398575</p>
                                        </div>
                                    </li>
                                </ul>
                            </div>
                        </div>
                        <div className="col-lg-8">
                            <div id="form_result">
                            </div>
                            <form id="contactpage" method="POST" className="contact-form wow slideInRight text-lg-left text-center">
                                <div className="row">
                                    <div className="col-lg-6 col-md-6">
                                        <div className="form-group mb-0">
                                            <input type="text" placeholder="Name" name="name" id="name" autoComplete="off" required />
                                        </div>
                                    </div>
                                    <div className="col-lg-6 col-md-6">
                                        <div className="form-group mb-0">
                                            <input type="email" id="emailHelp" name="emailHelp" placeholder="Email" autoComplete="off" required />
                                            <small className="form-text text-muted" />
                                        </div>
                                    </div>
                                    <div className="col-lg-6 col-md-6">
                                        <div className="form-group mb-0">
                                            <input type="tel" placeholder="Phone" name="phone" id="phone" required />
                                        </div>
                                    </div>
                                    <div className="col-lg-6 col-md-6">
                                        <div className="form-group mb-0">
                                            <input type="text" name="subject" placeholder="Subject" id="subject" />
                                        </div>
                                    </div>
                                </div>
                                <div className="row">
                                    <div className="col-lg-12">
                                        <div className=" form-group mb-0">
                                            <textarea placeholder="Message" rows={3} name="comments" id="comments" defaultValue={""} />
                                        </div>
                                    </div>
                                </div>
                                <button type="submit" id="submit" className="appointment-btn">Submit</button>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        </section>

    )
}

export default Contact