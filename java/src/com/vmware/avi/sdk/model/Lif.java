/*
 * Copyright 2021 VMware, Inc.
 * SPDX-License-Identifier: Apache License 2.0
 */

package com.vmware.avi.sdk.model;

import java.util.*;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;

/**
 * The Lif is a POJO class extends AviRestResource that used for creating
 * Lif.
 *
 * @version 1.0
 * @since 
 *
 */
@JsonIgnoreProperties(ignoreUnknown = true)
@JsonInclude(JsonInclude.Include.NON_NULL)
public class Lif  {
    @JsonProperty("cifs")
    private List<Cif> cifs;

    @JsonProperty("lif")
    private String lif;

    @JsonProperty("lif_label")
    private String lifLabel;

    @JsonProperty("subnet")
    private String subnet;


    /**
     * This is the getter method this will return the attribute value.
     * Field deprecated in 21.1.1.
     * @return cifs
     */
    public List<Cif> getCifs() {
        return cifs;
    }

    /**
     * This is the setter method. this will set the cifs
     * Field deprecated in 21.1.1.
     * @return cifs
     */
    public void setCifs(List<Cif>  cifs) {
        this.cifs = cifs;
    }

    /**
     * This is the setter method this will set the cifs
     * Field deprecated in 21.1.1.
     * @return cifs
     */
    public Lif addCifsItem(Cif cifsItem) {
      if (this.cifs == null) {
        this.cifs = new ArrayList<Cif>();
      }
      this.cifs.add(cifsItem);
      return this;
    }

    /**
     * This is the getter method this will return the attribute value.
     * Field deprecated in 21.1.1.
     * @return lif
     */
    public String getLif() {
        return lif;
    }

    /**
     * This is the setter method to the attribute.
     * Field deprecated in 21.1.1.
     * @param lif set the lif.
     */
    public void setLif(String  lif) {
        this.lif = lif;
    }

    /**
     * This is the getter method this will return the attribute value.
     * Field deprecated in 21.1.1.
     * @return lifLabel
     */
    public String getLifLabel() {
        return lifLabel;
    }

    /**
     * This is the setter method to the attribute.
     * Field deprecated in 21.1.1.
     * @param lifLabel set the lifLabel.
     */
    public void setLifLabel(String  lifLabel) {
        this.lifLabel = lifLabel;
    }

    /**
     * This is the getter method this will return the attribute value.
     * Field deprecated in 21.1.1.
     * @return subnet
     */
    public String getSubnet() {
        return subnet;
    }

    /**
     * This is the setter method to the attribute.
     * Field deprecated in 21.1.1.
     * @param subnet set the subnet.
     */
    public void setSubnet(String  subnet) {
        this.subnet = subnet;
    }


    @Override
    public boolean equals(java.lang.Object o) {
      if (this == o) {
          return true;
      }
      if (o == null || getClass() != o.getClass()) {
          return false;
      }
      Lif objLif = (Lif) o;
      return   Objects.equals(this.lif, objLif.lif)&&
  Objects.equals(this.lifLabel, objLif.lifLabel)&&
  Objects.equals(this.subnet, objLif.subnet)&&
  Objects.equals(this.cifs, objLif.cifs);
    }

    @Override
    public String toString() {
      StringBuilder sb = new StringBuilder();
      sb.append("class Lif {\n");
                  sb.append("    cifs: ").append(toIndentedString(cifs)).append("\n");
                        sb.append("    lif: ").append(toIndentedString(lif)).append("\n");
                        sb.append("    lifLabel: ").append(toIndentedString(lifLabel)).append("\n");
                        sb.append("    subnet: ").append(toIndentedString(subnet)).append("\n");
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
