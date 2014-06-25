/*
 * Copyright (c) 2014 Google Inc.
 * 
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * 
 * http://www.apache.org/licenses/LICENSE-2.0
 * 
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 */

package com.google.developers.actions.debugger.util;

import android.content.Context;
import android.util.Pair;

import java.util.List;

/**
 * Helper class to handle Actions.
 */
public class ActionData {
    private String mActionType;
    private String mAction;
    // TODO: hold a list of names, values instead.
    private List<Pair<String, String>> mActionExtras;

    public String getAction() {
        return mAction;
    }

    public void setAction(String action) {
        mAction = action;
    }

    public String getActionType() {
        return mActionType;
    }

    public void setActionType(String actionType) {
        mActionType = actionType;
    }

    public List<Pair<String, String>> getActionExtras() {
        return mActionExtras;
    }

    public void setActionExtras(List<Pair<String, String>> actionExtras) {
        mActionExtras = actionExtras;
    }

    public int getThumbnail(Context context) {
        String actionType = getActionType().toLowerCase();

        return context.getResources().getIdentifier(actionType, "drawable",
                context.getPackageName());
    }
}
