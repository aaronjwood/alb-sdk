/*
 * Copyright 2021 VMware, Inc.
 * SPDX-License-Identifier: Apache License 2.0
 */

package com.vmware.avi.sdk.model;

import java.util.Objects;
import com.fasterxml.jackson.annotation.JsonProperty;
import io.swagger.v3.oas.annotations.media.Schema;
import java.util.ArrayList;
import java.util.List;

/**
 * GslbGeoDbProfileApiResponse
 */
public class GslbGeoDbProfileApiResponse {
  @JsonProperty("count")
  private Integer count = null;

  @JsonProperty("results")
  private List<GslbGeoDbProfile> results = new ArrayList<GslbGeoDbProfile>();

  public GslbGeoDbProfileApiResponse count(Integer count) {
    this.count = count;
    return this;
  }

  /**
   * Get count
   * @return count
  **/
  @Schema(required = true, description = "")
  public Integer getCount() {
    return count;
  }

  public void setCount(Integer count) {
    this.count = count;
  }

  public GslbGeoDbProfileApiResponse results(List<GslbGeoDbProfile> results) {
    this.results = results;
    return this;
  }

  public GslbGeoDbProfileApiResponse addResultsItem(GslbGeoDbProfile resultsItem) {
    this.results.add(resultsItem);
    return this;
  }

  /**
   * Get results
   * @return results
  **/
  @Schema(required = true, description = "")
  public List<GslbGeoDbProfile> getResults() {
    return results;
  }

  public void setResults(List<GslbGeoDbProfile> results) {
    this.results = results;
  }

  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GslbGeoDbProfileApiResponse gslbGeoDbProfileApiResponse = (GslbGeoDbProfileApiResponse) o;
    return Objects.equals(this.count, gslbGeoDbProfileApiResponse.count) &&
        Objects.equals(this.results, gslbGeoDbProfileApiResponse.results);
  }

  @Override
  public int hashCode() {
    return Objects.hash(count, results);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GslbGeoDbProfileApiResponse {\n");
    
    sb.append("    count: ").append(toIndentedString(count)).append("\n");
    sb.append("    results: ").append(toIndentedString(results)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
}


