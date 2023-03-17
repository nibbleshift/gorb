<template>
    <v-container  fluid>
      <v-responsive class="pa-3 align-center  ">
        <div class="text-h4 pb-3">Benchmarks</div>
        <v-card elevation="1" class="mb-5">
	  <v-data-table
            show-expand
            :headers="headers" 
            :items="benches.edges">
	    <template v-slot:top>
              <v-toolbar class="bg-deep-purple">
                <v-toolbar-title>Benchmarks</v-toolbar-title>
              </v-toolbar>
            </template>
            <template v-slot:expanded-row="{ columns, item }">
              <tr>
                <td :colspan="columns.length">
                  <v-table>
                    <thead>
                      <tr>
                        <th>Name</th>
                        <th>Iterations</th>
                        <th>ns/op</th>
                        <th>alloced bytes/sec</th>
                        <th>allocs per op</th>
                        <th>MB/sec</th>
                        <th>Measured</th>
                        <th>Ord</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr v-for="r in item.value.node.benchResult" :key="Name">
                        <td>{{r.name}}</td>
                        <td>{{r.n}}</td>
                        <td>{{r.nsPerOp}}</td>
                        <td>{{r.allocedBytesPerOp}}</td>
                        <td>{{r.allocsPerOp}}</td>
                        <td>{{r.mbPerS}}</td>
                        <td>{{r.measured}}</td>
                        <td>{{r.ord}}</td>
                      </tr>
                    </tbody>
                  </v-table>
                </td>
              </tr>
            </template>
	  </v-data-table>
	</v-card>
      </v-responsive>
    </v-container>
</template>

<script>
import { gql } from "apollo-boost";

const FETCH_BENCHMARKS = gql`
        query benches {
          benches {
            edges {
              node {
                id
                os
                arch
                cpu
                package
                benchResult {
                  name
                  allocsPerOp
                  allocedBytesPerOp
                  mbPerS
                  n
                  nsPerOp
                }
              }
            }
          }
        }`;
export default {
  data() {
    return {
      benches: [],
      headers: [
        {
	  title: "Package",
	  key: "node.package",
	},
        {
	  title: "OS",
	  key: "node.os",
	},
        {
	  title: "Arch",
	  key: "node.arch",
	},
        {
	  title: "CPU",
	  key: "node.cpu",
	},
      ]
    }
  },
  apollo: {
    benches: {
      query: FETCH_BENCHMARKS,
    }
  }
};
</script>
