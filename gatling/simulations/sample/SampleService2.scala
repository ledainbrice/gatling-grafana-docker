package sample

import scala.concurrent.duration._
import io.gatling.core.Predef._
import io.gatling.http.Predef._
import io.gatling.jdbc.Predef._
import org.slf4j.LoggerFactory
import ch.qos.logback.classic.{Level, LoggerContext}

class SampleService2 extends Simulation {
	val httpProtocol = http
		.baseURL("http://sample-service2:4000")
		.acceptHeader("*/*")

	val scn = scenario("Sample Service Root 2")
		.exec(session => {
			println("Request: \n")
			session
		})
		.exec(http("request_1")
			.get("/ping")
			.check(status.is(200))
			.check(bodyString.saveAs("BODY"))) 

		.exec(session => {
			val response = session("BODY").as[String]
			println(s"Response body: \n$response")
			session
		})

	setUp(
    scn.inject(
			nothingFor(2 seconds)
      // ,atOnceUsers(10)
			// ,nothingFor(4 seconds)
      , rampUsers(10) over(10 seconds)
			, nothingFor(5 seconds)
      // ,rampUsers(100) over(10 seconds)
			// ,nothingFor(5 seconds)
     , rampUsersPerSec(30) to 300 during(10 minutes)
    )
  ).protocols(httpProtocol)
}
