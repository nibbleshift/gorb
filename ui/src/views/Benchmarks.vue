<template>
    <v-container  fluid>
      <v-responsive class="pa-3 align-center  ">
        <div class="text-h4 pb-3">Benchmarks</div>
        <v-card elevation="1">
          <v-table>
            <thead>
              <tr>
                <th class="text-left">
                  Package 
                </th>
                <th class="text-left">
                  OS
                </th>
                <th class="text-left">
                  Arch
                </th>
                <th class="text-left">
                  CPU
                </th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in benches.edges" :key="item.id">
                <td>{{ item.node.package}}</td>
                <td>{{ item.node.os }}</td>
                <td>{{ item.node.arch }}</td>
                <td>{{ item.node.cpu}}</td>
              </tr>
            </tbody>
          </v-table>
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
      benches: []
    }
  },
  apollo: {
    benches: {
      query: FETCH_BENCHMARKS,
    }
  }
};
</script>
