import PortfolioCard from "../components/PortfolioCard"

const Portfolio = () => {
    return (
        <>
            <section className="w-100 float-left portfolio-con padding-top" id="Portfolio">
                <div className="container">
                    <div className="generic-title text-center">
                        <h6 className="text-white">Creative Works</h6>
                        <h2 className="mb-0 text-white">Check My Portfolio</h2>
                    </div>
                    <div id="myBtnContainer" className="text-center">
                        <button className=" active active_button"> All </button>
                        <button>Products </button>
                        <button>Web App</button>
                        <button> Inetraction </button>
                        <button>Brand Identity</button>
                    </div>
                </div>
            </section>

            <PortfolioCard />
        </>
    )
}

export default Portfolio