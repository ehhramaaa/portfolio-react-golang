import Navbar from "../components/Navbar"
import Hero from "../components/Hero"

const Home = () => {
    return (
        <>
            <div className="header-and-banner-con w-100 float-left position-relative">
                <div className="header-and-banner-inner-con">
                    <header className="main-header">
                        <Navbar />
                    </header>
                    <section className="banner-main-con" id="home">
                        <div className="container pl-0 pr-0">
                            <Hero />
                        </div>
                    </section>
                </div>
            </div>
        </>
    )
}

export default Home