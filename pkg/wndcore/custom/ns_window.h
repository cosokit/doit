#ifndef NS_WINDOW_H
#define NS_WINDOW_H

#ifdef COCOA_UIKIT
#include <Cocoa/Cocoa.h>

void SetHasShadow(void *self, _Bool hasShadow) {
    NSWindow *window = self;
    [window setHasShadow: hasShadow];
}

#else
void SetHasShadow(void *self, _Bool hasShadow) {

}

#endif /* COCOA_UIKIT */

#endif /* WEBVIEW_H */
