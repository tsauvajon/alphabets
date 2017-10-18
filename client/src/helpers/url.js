import axios from 'axios'

const baseUrl = 'http://localhost:3333'

const apiPrefix = 'api'
const soccerPrefix = 'soccer'
const blockchainPrefix = 'blockchain'

const getTeamByIdFactory = (base) => (id) => axios.get(`${base}/teams/${id}`)

const getFixturesFactory = (base) => (begin, end) => axios.get(`${base}/fixtures/${begin}/${end}`)

const getOddsByFixtureFactory = (base) => (id) => axios.get(`${base}/odds/fixture/${id}`)

const getStandingsBySeasonIdFactory = (base) => (id) => axios.get(`${base}/standings/season/live/${id}`)

const test = (base) => () => axios.get(`${base}/test`)

const soccerFactory = (base) => ({
  getTeamById: getTeamByIdFactory(`${base}/${soccerPrefix}`),
  getFixtures: getFixturesFactory(`${base}/${soccerPrefix}`),
  getOddsByFixture: getOddsByFixtureFactory(`${base}/${soccerPrefix}`),
  getStandingsBySeasonId: getStandingsBySeasonIdFactory(`${base}/${soccerPrefix}`)
})

const blockchainFactory = (base) => ({
  test: test(`${base}/${blockchainPrefix}`)
})

const api = (baseUrl) => ({
  soccer: soccerFactory(`${baseUrl}/${apiPrefix}`),
  blockchain: blockchainFactory(`${baseUrl}/${apiPrefix}`)
})

export default api(baseUrl)
