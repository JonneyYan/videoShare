  <script>
  import {
    Line,
  } from 'vue-chartjs';

  export default Line.extend({
    name: 'LineChart',
    props: ['data'],
    data() {
      return {
        datacollection: this.data,
      };
    },
    methods: {
      rerender() {
        this.renderChart(this.datacollection, {
          responsive: true,
          maintainAspectRatio: false,
          fill: false,
        });
      },
    },
    mounted() {
      this.rerender();
    },
    updated() {
      this.rerender();
    },
  });
</script>
