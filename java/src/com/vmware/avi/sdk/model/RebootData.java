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
 * The RebootData is a POJO class extends AviRestResource that used for creating
 * RebootData.
 *
 * @version 1.0
 * @since 
 *
 */
@JsonIgnoreProperties(ignoreUnknown = true)
@JsonInclude(JsonInclude.Include.NON_NULL)
public class RebootData  {
    @JsonProperty("patch_version")
    private String patchVersion = null;

    @JsonProperty("reboot")
    private Boolean reboot = null;



    /**
     * This is the getter method this will return the attribute value.
     * Patch version for which reboot flag need to be computed.
     * Field introduced in 18.2.8, 20.1.1.
     * Default value when not specified in API or module is interpreted by Avi Controller as null.
     * @return patchVersion
     */
    public String getPatchVersion() {
        return patchVersion;
    }

    /**
     * This is the setter method to the attribute.
     * Patch version for which reboot flag need to be computed.
     * Field introduced in 18.2.8, 20.1.1.
     * Default value when not specified in API or module is interpreted by Avi Controller as null.
     * @param patchVersion set the patchVersion.
     */
    public void setPatchVersion(String  patchVersion) {
        this.patchVersion = patchVersion;
    }

    /**
     * This is the getter method this will return the attribute value.
     * This variable tells whether reboot has to be performed.
     * Field introduced in 18.2.8, 20.1.1.
     * Default value when not specified in API or module is interpreted by Avi Controller as null.
     * @return reboot
     */
    public Boolean getReboot() {
        return reboot;
    }

    /**
     * This is the setter method to the attribute.
     * This variable tells whether reboot has to be performed.
     * Field introduced in 18.2.8, 20.1.1.
     * Default value when not specified in API or module is interpreted by Avi Controller as null.
     * @param reboot set the reboot.
     */
    public void setReboot(Boolean  reboot) {
        this.reboot = reboot;
    }


    @Override
    public boolean equals(java.lang.Object o) {
      if (this == o) {
          return true;
      }
      if (o == null || getClass() != o.getClass()) {
          return false;
      }
      RebootData objRebootData = (RebootData) o;
      return   Objects.equals(this.patchVersion, objRebootData.patchVersion)&&
  Objects.equals(this.reboot, objRebootData.reboot);
    }

    @Override
    public String toString() {
      StringBuilder sb = new StringBuilder();
      sb.append("class RebootData {\n");
                  sb.append("    patchVersion: ").append(toIndentedString(patchVersion)).append("\n");
                        sb.append("    reboot: ").append(toIndentedString(reboot)).append("\n");
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
