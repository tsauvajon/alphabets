import axios from 'axios'

const baseUrl = 'http://localhost:3333'

const apiPrefix = 'api'
const soccerPrefix = 'soccer'
const blockchainPrefix = 'blockchain'

const getTeamById = (base) => (id) => axios.get(`${base}/teams/${id}`)

const getFixtures = (base) => (begin, end) => axios.get(`${base}/fixtures/${begin}/${end}`)
const getFixtureById = (base) => (id) => axios.get(`${base}/fixtures/${id}`)

const getHighlightsByFixtureFactory = (base) => (id) => axios.get(`${base}/highlights/fixture/${id}`)
const getOddsByFixture = (base) => (id) => axios.get(`${base}/odds/fixture/${id}`)

const test = (base) => () => axios.get(`${base}/test`)

const soccer = (base) => ({
  getTeamById: getTeamById(`${base}/${soccerPrefix}`),
  getFixtures: getFixtures(`${base}/${soccerPrefix}`),
  getFixtureById: getFixtureById(`${base}/${soccerPrefix}`),
  getOddsByFixture: getOddsByFixture(`${base}/${soccerPrefix}`),
  getHighlightsByFixture: getHighlightsByFixtureFactory(`${base}/${soccerPrefix}`)
})

const blockchain = (base) => ({
  test: test(`${base}/${blockchainPrefix}`)
})

const api = (baseUrl) => ({
  soccer: soccer(`${baseUrl}/${apiPrefix}`),
  blockchain: blockchain(`${baseUrl}/${apiPrefix}`)
})

export default api(baseUrl)
