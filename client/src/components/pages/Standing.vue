 <template>
   <v-data-table
       style="padding:5em;"
       :headers="headers"
       :items="standings"
       hide-actions
       class="elevation-1">
     <template slot="items" scope="props">
       <td>{{ props.item.position }}</td>
      <td class="text-xs-left"> <img :src="props.item.team_logo" style="height:100%; width:auto;"/></td>
       <td  class="text-xs-left"><span width="20%">{{ props.item.team_name }}</span></td>
       <td class="text-xs-right">{{ props.item.points }}</td>
       <td class="text-xs-right">{{ props.item.goals }}</td>
       <td class="text-xs-right">{{ props.item.wins }}</td>
       <td class="text-xs-right">{{ props.item.draws }}</td>
       <td class="text-xs-right">{{ props.item.lost }}</td>
       <td class="text-xs-right">{{ props.item.played }}</td>
       <td class="text-xs-right">{{ props.item.goal_diff }}</td>
     </template>
   </v-data-table>
 </template>


<script>
import { mapGetters } from 'vuex'
import Logo from '@/components/Logo'
import api from '@/helpers/url'

export default {
  name: 'standings',
  components: {
    Logo
  },

  data: () => ({
    headers:
    [
      {
        text: 'Position',

        align: 'left',
        value: 'Position'
      },
      {
        text: 'Logo',
        sortable: false,
        align: 'left',
        value: 'TeamLogo'
      },
      {
        text: 'Name',

        sortable: false,
        align: 'left',
        value: 'TeamName'
      },
      {
        text: 'Points',
        value: 'Points'
      },
      {
        text: 'Goals',
        value: 'Goals'
      },
      {
        text: 'Wins',
        value: 'Wins'
      },
      {
        text: 'Draws',
        value: 'Draws'
      },
      {
        text: 'Lost',
        value: 'Lost'
      },
      {
        text: 'Played',
        value: 'Played'
      },
      {
        text: 'Diff',
        value: 'Points'
      }
    ]
  }),

  computed: mapGetters(['standings']),

  async created () {
    const { saisonid } = this.$route.params
    console.log(saisonid)
    if (!saisonid || saisonid === 0) {
      this.$router.push('/')
    }

    if (!this.standings) {
      console.log('standings not loaded yet')
    }

    let standingsAsync
    console.log(api.soccer)

    try {
      standingsAsync = api.soccer.getStandingsBySeasonId(saisonid)
    } catch (ex) {
      // The request failed. Handle failures
      console.error(ex)
      return
    }

    const standings = await standingsAsync

    if (standings.status !== 200) {
      return
    }
    console.log(standings)
    this.$store.dispatch('setStandings', { standings: standings.data })
  }
}
</script>
 